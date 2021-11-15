package respModel

import "time"

type SumResp struct {
	LenUser      int `json:"users"`
	LenSelection int `json:"selections"`
	LenCover     int `json:"covers"`
}

// 搜索用户信息返回
type UserResp struct {
	Id        int    `json:"user_id"`
	Avatar    string `json:"avatar"`
	Nickname  string `json:"nickname"`
	Signature string `json:"slogan"`
}

// 搜索点歌信息返回
type SelectionResp struct {
	Id         int       `json:"selection_id"`
	Avatar     string    `json:"avatar"`
	Created_at time.Time `json:"post_time"`
	Song_name  string    `json:"song_name"`
	Nickname   string    `json:"nickname"`
}

// 搜索翻唱信息返回
type CoversResp struct {
	Id          int       `json:"cover_id"`
	Avatar      string    `json:"avatar"`
	Created_at  time.Time `json:"post_time"`
	Song_name   string    `json:"song_name"`
	Nickname    string    `json:"nickname"`
	Module      int       `json:"module"`
	SelectionId string    `json:"selection_id"`
	ClassicId   int       `json:"classic_id"`
}
