package models

import (
	"time"

	"git.100steps.top/100steps/healing2021_be/models/statements"
)

func TableInit() {
	go statements.AdvertisementInit()
	go statements.ClassicInit()
	go statements.CoverInit()
	go statements.PraiseInit()
	go statements.LotteryInit()
	go statements.MessageInit()
	go statements.MomentInit()
	go statements.MomentCommentInit()
	go statements.PrizeInit()
	go statements.SelectionInit()
	go statements.TaskInit()
	go statements.TaskTableInit()
	go statements.UserInit()
	time.Sleep(time.Second * 2)
}
