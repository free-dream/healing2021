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
	err := db.Model(&state.Selection{}).Where("id=?", SelectionId).Scan(&Module).Error
	return Module.Module, err
}

// 动态里面关于经典点歌和童年分享的区别处理，获取 song_id && moudle 参数
func DiffMoudle(SelectionId int, SongName string) (int, int, error) {
	module := 0
	songId := 0
	var err error
	if SelectionId != 0 {
		module, err = GetModuleBySelectionId(SelectionId)
		if err != nil {
			return 0, 0, err
		}

		if module == 2 {
			songId, err = GetClassicIdByName(SongName)
			if err != nil {
				return 0, 0, err
			}
		}
	}
	return module, songId, nil
}
