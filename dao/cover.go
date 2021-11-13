package dao

import (
	"fmt"
	//"errors"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/jinzhu/gorm"
)

// 获取所有翻唱信息（全表查找，后面优化
func GetCoverList(UserId int, ClassicId int) ([]respModel.CoverResp, error) {
	MysqlDB := setting.MysqlConn()
	var CoverResp []respModel.CoverResp
	var Cover []statements.Cover

	err := MysqlDB.Where("classic_id=?", ClassicId).Find(&Cover).Error
	if err != nil {
		fmt.Println(err)
		return CoverResp, err
	}

	// 参数转换
	for _, cover := range Cover {
		PlayerResp, err := GetPlayerInfo(UserId, int(cover.ID))
		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println(err)
			return CoverResp, err
		}

		coverResp := respModel.CoverResp{
			CoverId: int(cover.ID),
			Nickname: cover.Nickname,  // 翻唱者
			Avatar:   cover.Avatar,
			PostTime: tools.DecodeTime(cover.CreatedAt),
			File:PlayerResp.File,
			Name: PlayerResp.Name,  // 歌名
			Icon: PlayerResp.Icon,
			WorkName: PlayerResp.WorkName,
			Check: PlayerResp.Check,
		}
		CoverResp = append(CoverResp, coverResp)
	}
	return CoverResp, nil
}

// 通过翻唱歌曲id 获得 翻唱者id、翻唱歌名信息
type CoverInfo struct {
	Singer   int    `gorm:"user_id"`
	SongName string `gorm:"song_name"`
}

func GetCoverInfo(CoverId int) (int, string, error) {
	MysqlDB := setting.MysqlConn()
	coverInfo := CoverInfo{}
	err := MysqlDB.Table("cover").Select("user_id,song_name").Where("id=?", CoverId).Scan(&coverInfo).Error
	return coverInfo.Singer, coverInfo.SongName, err
}

// 判断某用户是否点赞
func HaveCoverLaud(UserId int, CoverId int) int {
	MysqlDB := setting.MysqlConn()
	err := MysqlDB.Where("user_id=? and cover_id=? and is_liked=?", UserId, CoverId, 1).First(&statements.Praise{}).Error
	if gorm.IsRecordNotFoundError(err) {
		return 0
	} else if err != nil {
		return -1
	}
	return 1
}
