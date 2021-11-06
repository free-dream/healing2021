package respModel

// 推荐童年原唱歌曲信息
type ClassicResp struct {
	ClassicId int    `json:"classic_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Click     int    `json:"click"`
}

// 童年歌曲列表翻唱信息
type ClassicListResp struct {
	ClassicId int    `json:"classic_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	WorkName  string `json:"work_name"`
}

// 童年原唱页原唱信息
type OriginInfoResp struct {
	ClassicId int    `json:"classic_id"`
	SongName  string `json:"song_name"`
	Singer    string `json:"singer"`
	Icon      string `json:"icon"`
	WorkName string `json:"work_name"`
}
