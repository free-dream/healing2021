package respModel

type CoverResp struct {
	CoverId  int    `json:"cover_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	PostTime string `json:"post_time"`
}

type HotSong struct {
	SongName string `json:"song_name"`
	Language string `json:"language"`
	Style    string `json:"style"`
}