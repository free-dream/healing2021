package respModel

type CoverResp struct {
	Nickname string    `json:"nickname,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	PostTime string `json:"post_time"`
}

type PlayerChildResp struct {
	SongName string `json:"song_name,omitempty"`
	File     string `json:"file,omitempty"`
	Lyrics   string `json:"lyrics,omitempty"`
	Icon     string `json:"icon,omitempty"`
	WorkName string `json:"work_name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}

type PlayerNormalResp struct {
	SongName string `json:"song_name,omitempty"`
	File     string `json:"file,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Nickname string `json:"nickname,omitempty"`
}