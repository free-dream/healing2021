package respModel

// 推荐童年歌曲信息
type ClassicResp struct {
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Click int    `json:"click"`
}

// 童年歌曲列表版本信息
type ClassicListResp struct {
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Time string `json:"time"`
	WorkName    string `json:"work_name"`
}

// 童年歌曲完全版信息
type OriginInfoResp struct {
	SongId int    `json:"song_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Singer string `json:"singer,omitempty"`
	Icon   string `json:"icon,omitempty"`
}