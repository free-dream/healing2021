package dao

import (
	//"errors"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/jinzhu/gorm"
)

// 获取所有翻唱信息（全表查找，后面优化)
func GetCoverList(UserId int, ClassicId int) ([]respModel.CoverResp, error) {
	MysqlDB := setting.MysqlConn()
	var coversResp []respModel.CoverResp
	var covers []statements.Cover

	err := MysqlDB.Where("classic_id=?", ClassicId).Find(&covers).Error
	if err != nil {
		return coversResp, err
	}

	// 参数转换
	for _, cover := range covers {
		workName, err := GetClassicSongFromCover(cover.SongName)
		if err != nil && err != gorm.ErrRecordNotFound {
			return coversResp, err
		}
		nickname, err := GetUserNickname(cover.UserId)
		if err != nil && err != gorm.ErrRecordNotFound {
			return coversResp, err
		}
		cResp := respModel.CoverResp{
			UserId:   cover.UserId,
			CoverId:  int(cover.ID),
			Nickname: nickname, // 翻唱者
			Avatar:   cover.Avatar,
			PostTime: tools.DecodeTime(cover.CreatedAt),
			File:     cover.File,
			Name:     cover.SongName, // 歌名
			Icon:     cover.Avatar,
			WorkName: workName,
			Check:    HaveCoverLaud(UserId, int(cover.ID)),
		}
		coversResp = append(coversResp, cResp)
	}
	return coversResp, nil
}

// 通过翻唱歌曲id 获得 翻唱者id、翻唱歌名信息
type CoverInfo struct {
	Singer   int    `gorm:"user_id"`
	SongName string `gorm:"song_name"`
}

func GetCoverInfo(CoverId int) (int, string, error) {
	MysqlDB := setting.MysqlConn()
	//coverInfo := CoverInfo{}
	//err := MysqlDB.Table("cover").Select("user_id,song_name").Where("id=?", CoverId).Scan(&coverInfo).Error
	coverInfo := statements.Cover{}
	err := MysqlDB.Where("id=?", CoverId).First(&coverInfo).Error
	if err != nil {
		return 0, "", err
	}
	//return coverInfo.Singer, coverInfo.SongName, err
	return coverInfo.UserId, coverInfo.SongName, err
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

// 通过 cover_id 找 童年歌名、来源
type ClassicWorkName struct {
	WorkName string `gorm:"work_name"`
}

func GetClassicSongFromCover(songName string) (string, error) {
	db := setting.MysqlConn()
	workName := ClassicWorkName{}
	err := db.Model(&statements.Classic{}).Select("work_name").Where("song_name=?", songName).Scan(&workName).Error
	if err != nil {
		return "", err
	}
	return workName.WorkName, nil
}
