package respModel

type CoverResp struct {
	CoverId  int    `json:"cover_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	PostTime string `json:"post_time"`
}
