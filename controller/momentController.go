package controller

import (
	"git.100steps.top/100steps/healing2021_be/controller/ws"
	"strconv"
	"time"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"git.100steps.top/100steps/healing2021_be/task"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 拉取广场动态列表[三种模式:new/recommend/search]
func GetMomentList(ctx *gin.Context) {
	var MomentsResp []respModel.MomentResp
	// 获取 url 参数
	Method := ctx.Param("method")
	Keyword := ctx.Query("keyword")
	page := ctx.Query("page")

	// 模式判断和处理
	if Method != "new" && Method != "recommend" && Method != "search" {
		ctx.JSON(403, e.ErrMsgResponse{Message: "模式选择出错"})
		return
	}

	// page 参数合法性判断
	Page, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "page参数非法"})
		return
	}

	if Keyword != "" {
		sandwich.PutInSearchWord(Keyword)
	}

	// 从数据库中得到经过筛选的一页 Momment 列表
	AllMoment, ok := dao.GetMomentPage(Method, Keyword, Page)
	if !ok {
		ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
		return
	}

	// 获取和整理其他所需信息，装进 response
	for _, OneMoment := range AllMoment {
		User := statements.User{}
		UserId := sessions.Default(ctx).Get("user_id").(int) // 获取当前用户 id
		User, ok := dao.GetUserById(OneMoment.UserId)
		if !ok {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库查询失败"})
			return
		}

		// 如有点歌，查表判断点歌类型 && 根据模式返回对应的 song_id
		module, songId, err := dao.DiffMoudle(OneMoment.SelectionId, OneMoment.SongName)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库查询失败"})
		}

		TmpMoment := respModel.MomentResp{
			Content:    OneMoment.Content,
			DynamicsId: int(OneMoment.ID),
			CreatedAt:  tools.DecodeTime(OneMoment.CreatedAt),
			Song:       OneMoment.SongName,
			SongId:     songId,
			Module:     module,
			Lauds:      dao.CountMLaudsById(int(OneMoment.ID)),
			Lauded:     dao.HaveMLauded(UserId, int(OneMoment.ID)),
			Comments:   dao.CountCommentsById(int(OneMoment.ID)),
			Status:     tools.DecodeStrArr(OneMoment.State),
			Creator:    dao.TransformUserInfo(User),
		}
		MomentsResp = append(MomentsResp, TmpMoment)
	}

	ctx.JSON(200, MomentsResp)
}

// 发布动态
// @@@@@@@任务模块已植入此接口函数@@@@@@@
type MomentBase struct {
	Content string   `json:"content"`
	Status  []string `json:"status"`

	HaveSong int `json:"have_selection"`

	SongName string `json:"song_name"`
	Language string `json:"language"`
	Style    string `json:"style"`
	Remark   string `json:"remark"`

	ClassicId int `json:"classic_id"`
}

func PostMoment(ctx *gin.Context) {
	// 参数绑定
	var NewMoment MomentBase
	var Moment statements.Moment
	ctx.ShouldBind(&NewMoment)

	param := statements.Selection{}
	//把userid拿出来用于任务---voloroloq 2021.11.1
	userid := sessions.Default(ctx).Get("user_id").(int)
	param.UserId = userid

	switch NewMoment.HaveSong {
	case 0:
		param.SongName = NewMoment.SongName
		param.Language = NewMoment.Language
		param.Remark = NewMoment.Remark
		param.Style = NewMoment.Style
		resp, err := dao.Select(param)
		if err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "点歌操作失败"})
			return
		}
		Moment.SelectionId = resp.ID
		Moment.Type = 0
	case 1:
		Moment.ClassicId = NewMoment.ClassicId
		Moment.Type = 1
	case 2:
		Moment.Type = 2
	default:// 出现错误
		ctx.JSON(403, e.ErrMsgResponse{Message: "非法参数"})
		//为了保证后面任务在接口使用时顺利进行---voloroloq 2021.11.1
		return
	}

	// 统计大家的状态、统计点歌情况
	for _, state := range NewMoment.Status {
		sandwich.PutInStates(state)
	}
	if NewMoment.HaveSong == 0 {
		sandwich.PutInHotSong(tools.EncodeSong(tools.HotSong{
			SongName: param.SongName,
			Language: param.Language,
			Style:    param.Style,
		}))
	}

	// 添加参数
	Moment.Content=NewMoment.Content
	Moment.SongName=NewMoment.SongName
	Moment.UserId=param.UserId
	Moment.State=tools.EncodeStrArr(NewMoment.Status)

	// 存入数据库
	if ok := dao.CreateMoment(Moment); !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
		return
	}

	//任务模块植入 2021.11.1
	thistask := task.MT
	thistask.AddRecord(userid)

	ctx.JSON(200, e.ErrMsgResponse{Message: "动态发布成功"})
}

// 查看动态的详情
func GetMomentDetail(ctx *gin.Context) {
	// url 参数的获取和合法性判断
	momentIdStr := ctx.Param("id")
	if momentIdStr == "" {
		ctx.JSON(403, e.ErrMsgResponse{Message: "动态id参数未传入"})
		return
	}
	Id, err := strconv.Atoi(momentIdStr)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "id参数非法"})
		return
	}

	// 数据库单条查找
	Moment, ok := dao.GetMomentById(Id)
	if !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库查找失败"})
		return
	}

	// 数据转换
	UserId := sessions.Default(ctx).Get("user_id").(int) // 获取当前用户 id
	User, ok_ := dao.GetUserById(Moment.UserId)
	if !ok_ {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库查询失败"})
		return
	}

	// 如有点歌或分享，查表判断点歌类型 && 根据模式返回对应的 song_id
	module, songId, err := dao.DiffMoudle(Moment.SelectionId, Moment.SongName)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库查询失败"})
	}

	MomentDetail := respModel.MomentResp{
		DynamicsId: int(Moment.ID),
		Content:    Moment.Content,
		CreatedAt:  tools.DecodeTime(Moment.CreatedAt),
		Song:       Moment.SongName,
		SongId:     songId,
		Module:     module,
		Lauds:      dao.CountMLaudsById(int(Moment.ID)),
		Lauded:     dao.HaveMLauded(UserId, int(Moment.ID)),
		Comments:   dao.CountCommentsById(int(Moment.ID)),
		Status:     tools.DecodeStrArr(Moment.State),
		Creator:    dao.TransformUserInfo(User),
	}

	ctx.JSON(200, MomentDetail)
}

// 给动态添加评论
type CommentBase struct {
	DynamicsId int    `json:"dynamics_id"`
	Content    string `json:"content"`
}

func PostComment(ctx *gin.Context) {
	var NewComment CommentBase
	if err := ctx.ShouldBind(&NewComment); err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "评论参数不完整"})
		return
	}

	// 转换参数
	UserId := sessions.Default(ctx).Get("user_id").(int) // 获取当前用户 id
	Comment := statements.MomentComment{
		Comment:  NewComment.Content,
		MomentId: NewComment.DynamicsId,
		UserId:   UserId,
	}

	// 存入数据库
	commentId := 0
	ok := false
	if commentId, ok = dao.CreateComment(Comment); !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
		return
	}

	// 发送相应的系统消息[有 实际评论写入成功，但是系统消息发送失败 的不一致风险]
	conn := ws.GetConn()
	userId, err := dao.GetMomentSenderId(NewComment.DynamicsId)
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
		return
	}
	err = conn.SendSystemMsg(respModel.SysMsg{
		Uid:       uint(userId),
		Type:      3,
		ContentId: uint(commentId),
		Time:      time.Now(),
	})
	if err != nil {
		ctx.JSON(500, e.ErrMsgResponse{Message: "系统消息发送失败"})
		return
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "评论发布成功"})
}

// 拉取动态的评论列表
func GetCommentList(ctx *gin.Context) {
	var CommentsResp []respModel.CommentResp

	// 获取 url 参数并判断合法性
	CommentIdstr := ctx.Param("id")
	if CommentIdstr == "" {
		ctx.JSON(403, e.ErrMsgResponse{Message: "获取动态ID失败"})
		return
	}
	CommentId, err := strconv.Atoi(CommentIdstr)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "ID参数非法"})
		return
	}

	// 数据库筛选
	CommentsList, ok := dao.GetCommentsByMomentId(CommentId)
	if !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库操作失败或评论为空"})
		return
	}

	// 参数转换，填入响应
	UserId := sessions.Default(ctx).Get("user_id").(int) // 获取当前用户 id
	for _, comment := range CommentsList {
		User, ok := dao.GetUserById(comment.UserId)
		if !ok {
			ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
			return
		}

		Comment := respModel.CommentResp{
			CommentId: int(comment.ID),
			Content:   comment.Comment,
			Lauds:     dao.CountCLaudsById(int(comment.ID)),
			Lauded:    dao.HaveCLauded(UserId, int(comment.ID)),
			Creator:   dao.TransformUserInfo(User),
			CreatedAt: tools.DecodeTime(comment.CreatedAt),
		}
		CommentsResp = append(CommentsResp, Comment)
	}

	ctx.JSON(200, CommentsResp)
}

// 动态搜索推荐
func DynamicsSearchHot(ctx *gin.Context) {
	result := sandwich.GetSearchWord()
	ctx.JSON(200, result)
}

// 大家的状态推荐
func OursStates(ctx *gin.Context) {
	result := sandwich.GetStates()
	ctx.JSON(200, result)
}

// 点歌页热门歌曲推荐推荐
func HotSong(ctx *gin.Context) {
	result := sandwich.GetHotSong()

	var hotSongResp []tools.HotSong
	for _, hotSong := range result {
		hotSongResp = append(hotSongResp, tools.DecodeSong(hotSong))
	}

	ctx.JSON(200, hotSongResp)
}
