package dao

import (
	state "git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

type SelectionModule struct {
	Module int `gorm:"module"`
}

func GetModuleBySelectionId(SelectionId int) (int, error) {
	db := setting.MysqlConn()
	Module := SelectionModule{}
	ok := db.Model(&state.Selection{}).Where("id=?", SelectionId).Scan(&Module).Error
	return Module.Module, ok
}
