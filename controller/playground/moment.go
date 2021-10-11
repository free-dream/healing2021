package playground

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
	"time"
)

// 用户信息
type User struct {
	Id int
	Nackname string
	Avatar string
	Avatar_visible int
}

// 拉取广场动态列表[三种模式:new/recommend/search]
type MomentResp struct {
	DynamicsId int    `json:"dynamics_id"`
	Content    string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Img       []string  `json:"img"`
	Song        string    `json:"song"`
	Lauds       int       `json:"lauds"`
	Lauded      int       `json:"lauded"`
	Comments    int       `json:"comments"`
	Status      []string  `json:"status"`
	Creator     User      `json:"creator"`
}
func GetMomentList(ctx *gin.Context) {
	var MomentsResp []MomentResp
	var TmpMoment  MomentResp
	// 获取 url 参数
	method := ctx.Param("method")
	keyword := ctx.Query("keyword")

	// 模式判断和处理
	if method != "new" && method != "recommend" && method != "search"{
		ctx.JSON(403, e.ErrMsgResponse{Message: "模式选择出错"})
		return
	}

	// 从数据库中得到经过筛选的 Momment 列表
	AllMoment, ok := dao.GetAllMoment(method, keyword)
	if !ok {
		ctx.JSON(403, e.ErrMsgResponse{Message: "数据库查询失败"})
		return
	}

	// 获取和整理其他所需信息，装进 response
	for _, OneMoment := range AllMoment{
		TmpMoment.Content = OneMoment.Content
		TmpMoment.DynamicsId = int(OneMoment.ID)
		TmpMoment.CreatedAt = OneMoment.CreatedAt
		TmpMoment.Img = tools.DecodeStrArr(OneMoment.Picture)
		//TmpMoment.Song = dao.GetSongNameById(OneMoment.SongId)
		TmpMoment.Lauds = dao.CountLaudsById(TmpMoment.DynamicsId)
		TmpMoment.Lauded = dao.HaveLauded(TmpMoment.DynamicsId)
		TmpMoment.Comments = dao.CountCommentsById(TmpMoment.DynamicsId)
		TmpMoment.Status = tools.DecodeStrArr(OneMoment.States)
		//TmpMoment.Creator

		MomentsResp = append(MomentsResp, TmpMoment)
	}

	ctx.JSON(200, MomentsResp)
}

// 发布动态
type MomentBase struct {
	Content string
	Img []string
	Song string
	Status []string
}
func PostMoment(ctx *gin.Context) {
	var NewMoment MomentBase
	if err:= ctx.ShouldBind(&NewMoment); err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message:"动态发布失败"})
		return
	}

	// 存进数据库

	ctx.JSON(200, "")
}

// 查看动态的详情
func GetMomentDetail(ctx *gin.Context) {
	// url 参数的获取和合法性判断
	momentIdStr := ctx.Param("id")
	if momentIdStr == ""{
		ctx.JSON(403, e.ErrMsgResponse{Message: "该动态不存在"})
		return
	}

	var MomentDetail MomentResp

	// 数据库单条查找

	ctx.JSON(200, MomentDetail)
}

// 给动态添加评论
type CommentBase struct {
	Dynamics_id int    `json:"dynamics_id"`
	Content     string `json:"content"`
}
func PostComment(ctx *gin.Context) {
	var NewComment CommentBase
	if err:= ctx.ShouldBind(&NewComment); err != nil {
		ctx.JSON(403, e.ErrMsgResponse{Message:"评论发布失败"})
		return
	}

	// 将评论存进数据库

	ctx.JSON(200, "")
}

// 拉取动态的评论列表
type CommentResp struct {
	Creator    User      `json:"creator"`
	Comment_id int       `json:"comment_id"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_at"`
	Lauds      int       `json:"lauds"`
	Lauded     int       `json:"lauded"`
}
func GetCommentList(ctx *gin.Context) {
	var CommentsResp []CommentResp
	var Comment CommentResp

	// 获取 url 参数
	CommentIdstr := ctx.Param("id")
	if CommentIdstr == ""{
		ctx.JSON(403, e.ErrMsgResponse{Message: "获取评论ID失败"})
		return
	}

	// 数据库筛选
	// 填入响应

	ctx.JSON(200, CommentsResp)
}

// 给动态或评论点赞（取消点赞）
func PriseOrNot(ctx *gin.Context) {
	// url 参数获取
	Types := ctx.Param("type")
	Id := ctx.Param("id")

	// 写入数据库
	//ctx.JSON(403, e.ErrMsgResponse{Message: "点赞或取消点赞失败"})
	//return

	ctx.JSON(200, "")
}
