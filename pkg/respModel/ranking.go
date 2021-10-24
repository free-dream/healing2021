package respModel

import "time"

//排序的返回值
type RankingResp struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

//每日热榜返回值
type HotResp struct {
	Avatar   string    `json:"avatar"`
	Nickname string    `json:"nickname"`
	Posttime time.Time `json:"post_time"`
	Likes    int       `json:"likes"`
	Songname string    `json:"song_name"`
}
