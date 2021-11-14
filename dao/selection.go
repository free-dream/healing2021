package dao

import (
	state "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type SelectorRemark struct {
	Remark string `gorm:"remark"`
}

func GetSelectorInfo(selectionId int) (string, error) {
	db := setting.MysqlConn()
	remark := SelectorRemark{}
	err := db.Model(&state.Selection{}).Select("remark").Where("id=?", selectionId).Scan(&remark).Error
	if err != nil {
		return "", err
	}
	return remark.Remark, nil
}
