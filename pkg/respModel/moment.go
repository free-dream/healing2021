package respModel

// 用户信息
type UserInfo struct {
	Id            int    `json:"id"`
	Nackname      string `json:"nackname"`
	Avatar        string `json:"avatar"`
	AvatarVisible int    `json:"avatar_visible"`
}

// 动态响应
type MomentResp struct {
	DynamicsId int      `json:"dynamics_id"`
	Content    string   `json:"content"`
	CreatedAt  string   `json:"created_at"`
	Song       string   `json:"song"`
	SongId     int      `json:"song_id"`
	Module     int      `json:"module"`
	Lauds      int      `json:"lauds"`
	Lauded     int      `json:"lauded"`
	Comments   int      `json:"comments"`
	Status     []string `json:"status"`
	Creator    UserInfo `json:"creator"`
}

// 评论响应
type CommentResp struct {
	Creator   UserInfo `json:"creator"`
	CommentId int      `json:"comment_id"`
	Content   string   `json:"content"`
	CreatedAt string   `json:"created_at"`
	Lauds     int      `json:"lauds"`
	Lauded    int      `json:"lauded"`
}
