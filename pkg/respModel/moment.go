package respModel

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"time"
)

// 用户信息
type UserInfo struct {
	Id            int    `json:"id"`
	Nackname      string `json:"nackname"`
	Avatar        string `json:"avatar"`
	AvatarVisible int    `json:"avatar_visible"`
}

// 将数据库中的用户信息 User 进行提取转换为 UserInfo【未完工】
func TransformUserInfo(OneUser statements.User) UserInfo {
	var userInfo UserInfo
	return userInfo
}

// 动态响应
type MomentResp struct {
	DynamicsId  int       `json:"dynamics_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
	Song        string    `json:"song"`
	SelectionId int       `json:"selection_id"`
	Lauds       int       `json:"lauds"`
	Lauded      int       `json:"lauded"`
	Comments    int       `json:"comments"`
	Status      []string  `json:"status"`
	Creator     UserInfo  `json:"creator"`
}

// 评论响应
type CommentResp struct {
	Creator   UserInfo  `json:"creator"`
	CommentId int       `json:"comment_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Lauds     int       `json:"lauds"`
	Lauded    int       `json:"lauded"`
}
