package models

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
)

func TableInit() {
	statements.AdvertisementInit()
	statements.ClassicInit()
	statements.CoverInit()
	statements.PraiseInit()
	statements.LotteryInit()
	statements.MessageInit()
	statements.MomentInit()
	statements.MomentCommentInit()
	statements.PrizeInit()
	statements.SelectionInit()
	statements.TaskInit()
	statements.TaskTableInit()
	statements.UserInit()
}
