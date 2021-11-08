package respModel

//查询点赞表获得的视图
type CoverRank struct {
	CoverId int
	Likes   int
}

//排序的返回值
type RankingResp struct {
	Userid   int    `json:"userid"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

//当前用户的返回结构
type RankingUResp struct {
	Rank string `json:"rank"`
}

//每日热榜返回值
type HotResp struct {
	CoverId  int    `json:"cover_id"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Posttime string `json:"post_time"`
	Likes    int    `json:"likes"`
	Songname string `json:"song_name"`
}
