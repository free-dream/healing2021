package respModel

// 歌曲页的响应
type PlayerResp struct {
	CoverId  int    `json:"cover_id"`
	File     string `json:"file"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Icon     string `json:"icon"`
	WorkName string `json:"work_name"`
}
