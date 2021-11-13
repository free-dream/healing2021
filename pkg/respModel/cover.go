package respModel

type CoverResp struct {
	CoverId  int    `json:"cover_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	PostTime string `json:"post_time"`
	File     string `json:"file"`
	Name     string `json:"name"` // 歌曲名
	Icon     string `json:"icon"`
	WorkName string `json:"work_name"`
	Check    int    `json:"check"`
}
