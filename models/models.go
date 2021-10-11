package models

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
)

func TableInit() {
	statements.AdvertisementInit()
	statements.ClassicInit()
	statements.CoverInit()
	statements.CoverLikeInit()
	statements.LotteryInit()
	statements.MessageInit()
	statements.MomentInit()
	statements.MomentLikeInit()
	statements.MomentCommentInit()
	statements.PrizeInit()
	statements.SelectionInit()
	statements.SongInit()
	statements.SongLikeInit()
	statements.TaskInit()
	statements.TaskTableInit()
	statements.UserInit()
}
