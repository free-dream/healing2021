package playground

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 拉取广场动态列表[三种模式:new/recommend/search]
func GetMomentList(ctx *gin.Context) {
	var MomentsResp []respModel.MomentResp
	var TmpMoment respModel.MomentResp
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

	// 从数据库中得到经过筛选的一页 Momment 列表
	AllMoment, ok := dao.GetMomentPage(Method, Keyword, Page)
	if !ok {
		ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
		return
	}

	// 获取和整理其他所需信息，装进 response
	for _, OneMoment := range AllMoment {
		// 错误判断还没做
		TmpMoment.Content = OneMoment.Content
		TmpMoment.DynamicsId = int(OneMoment.ID)
		TmpMoment.CreatedAt = OneMoment.CreatedAt
		TmpMoment.Song = OneMoment.SongName
		TmpMoment.SelectionId = OneMoment.SelectionId
		TmpMoment.Lauds = dao.CountMLaudsById(TmpMoment.DynamicsId)
		//把缓存加进来之前注释的部分都用不了
		UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
		TmpMoment.Lauded = dao.HaveMLauded(int(UserId), TmpMoment.DynamicsId)

		TmpMoment.Comments = dao.CountCommentsById(TmpMoment.DynamicsId)
		TmpMoment.Status = tools.DecodeStrArr(OneMoment.State)
		User := statements.User{}
		User, ok := dao.GetUserById(OneMoment.UserId)
		fmt.Println(User)
		fmt.Println(OneMoment.UserId)
		if !ok {
			ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
			return
		}
		TmpMoment.Creator = respModel.TransformUserInfo(User)
		MomentsResp = append(MomentsResp, TmpMoment)
	}

	ctx.JSON(200, MomentsResp)
}

// 发布动态
type MomentBase struct {
	Content     string   `json:"content"`
	Song        string   `json:"song"`
	Status      []string `json:"status"`
	SelectionId int      `json:"selection_id"`
}

func PostMoment(ctx *gin.Context) {
	// 参数绑定
	var NewMoment MomentBase
	ctx.ShouldBind(&NewMoment)

	// 转换参数
	var Moment statements.Moment
	Moment.Content = NewMoment.Content
	Moment.SongName = NewMoment.Song
	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	Moment.UserId = int(UserId)
	Moment.State = tools.EncodeStrArr(NewMoment.Status)
	Moment.SelectionId = NewMoment.SelectionId

	// 存入数据库
	if ok := dao.CreateMoment(Moment); !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
		return
	}

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
	var MomentDetail respModel.MomentResp
	MomentDetail.DynamicsId = int(Moment.ID)
	MomentDetail.Content = Moment.Content
	MomentDetail.CreatedAt = Moment.CreatedAt
	MomentDetail.Song = Moment.SongName
	MomentDetail.Lauds = dao.CountMLaudsById(MomentDetail.DynamicsId)
	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	MomentDetail.Lauded = dao.HaveMLauded(int(UserId), MomentDetail.DynamicsId)
	MomentDetail.Comments = dao.CountCommentsById(MomentDetail.DynamicsId)
	MomentDetail.Status = tools.DecodeStrArr(Moment.State)
	User, ok_ := dao.GetUserById(Moment.UserId)
	if !ok_ {
		ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
		return
	}
	MomentDetail.Creator = respModel.TransformUserInfo(User)

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
	var Comment statements.MomentComment
	Comment.Comment = NewComment.Content
	Comment.MomentId = NewComment.DynamicsId
	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	Comment.UserId = int(UserId)

	// 存入数据库
	if ok := dao.CreateComment(Comment); !ok {
		ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
		return
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "评论发布成功"})
}

// 拉取动态的评论列表
func GetCommentList(ctx *gin.Context) {
	var CommentsResp []respModel.CommentResp
	var Comment respModel.CommentResp

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
	for _, comment := range CommentsList {
		Comment.CommentId = int(comment.ID)
		Comment.Content = comment.Comment
		UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
		Comment.Lauded = dao.HaveCLauded(int(UserId), Comment.CommentId)
		Comment.Lauds = dao.CountCLaudsById(Comment.CommentId)

		User, ok := dao.GetUserById(comment.UserId)
		if !ok {
			ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
			return
		}
		Comment.Creator = respModel.TransformUserInfo(User)
		Comment.CreatedAt = comment.CreatedAt

		CommentsResp = append(CommentsResp, Comment)
	}

	ctx.JSON(200, CommentsResp)
}

// 给动态或评论点赞（取消点赞）
func PriseOrNot(ctx *gin.Context) {
	// url 参数获取
	Types := ctx.Param("type")
	Idstr := ctx.Param("id")
	Id, err := strconv.Atoi(Idstr)
	if err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message: "id参数非法"})
		return
	}

	UserId := tools.GetUser(ctx.Copy()).ID // 获取当前用户 id
	// 分模式进行点赞处理
	if Types == "comment" {
		if err := dao.CLaudedById(Id, int(UserId)); err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
			return
		}
	} else if Types == "moment" {
		if err := dao.MLaudedById(Id, int(UserId)); err != nil {
			ctx.JSON(500, e.ErrMsgResponse{Message: "数据库写入失败"})
			return
		}
	} else {
		ctx.JSON(403, e.ErrMsgResponse{Message: "type参数有误"})
		return
	}

	ctx.JSON(200, e.ErrMsgResponse{Message: "点赞或取消点赞成功"})
}
