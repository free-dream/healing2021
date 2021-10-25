package respModel

import "time"

type SumResp struct {
	LenUser      int `json:"users"`
	LenSelection int `json:"selections"`
	LenCover     int `json:"covers"`
}

// 搜索用户信息返回
type UserResp struct {
	Userid   int    `json:"user_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Slogan   string `json:"slogan"`
}

// 搜索点歌信息返回
type SelectionResp struct {
	Selectionid int       `json:"selection_id"`
	Avatar      string    `json:"avatar"`
	Posttime    time.Time `json:"post_time"`
	Songname    string    `json:"song_name"`
	Nickname    string    `json:"nickname"`
}

// 搜索翻唱信息返回
type CoversResp struct {
	Coverid  int       `json:"cover_id"`
	Avatar   string    `json:"avatar"`
	Posttime time.Time `json:"post_time"`
	Songname string    `json:"song_name"`
	Nickname string    `json:"nickname"`
}
