package dao

import (
	//"errors"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

// 获取所有翻唱信息（全表查找，后面优化
func GetCoverList(ClassicId int) ([]respModel.CoverResp, error) {
	MysqlDB := setting.MysqlConn()
	var CoverResp []respModel.CoverResp
	var Cover []statements.Cover

	err := MysqlDB.Where("classic_id=?", ClassicId).Find(&Cover).Error
	if err != nil {
		return CoverResp, err
	}

	// 参数转换
	for _, cover := range Cover {
		coverResp := respModel.CoverResp{
			CoverId:  cover.ClassicId,
			Nickname: cover.Nickname,
			Avatar:   cover.Avatar,
			PostTime: tools.DecodeTime(cover.CreatedAt),
		}
		CoverResp = append(CoverResp, coverResp)
	}
	return CoverResp, nil
}

// 通过翻唱歌曲id 获得 翻唱者id、翻唱歌名信息
type CoverInfo struct {
	Singer int `gorm:"user_id"`
	SongName string `gorm:"song_name"`
}
func GetCoverInfo(CoverId int)  (int,string,error){
	MysqlDB := setting.MysqlConn()
	coverInfo := CoverInfo{}
	err := MysqlDB.Where("id=?", CoverId).Scan(&coverInfo).Error
	return coverInfo.Singer, coverInfo.SongName, err
}
