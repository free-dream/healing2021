package dao

import (
	"errors"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

// 获取所有翻唱信息（全表查找，后面优化
func GetCoverList(Name string) ([]respModel.CoverResp, error) {
	MysqlDB := setting.MysqlConn()
	var CoverResp []respModel.CoverResp
	var Cover []statements.Cover

	err := MysqlDB.Where("nickname=?", Name).Find(&Cover).Error
	if err != nil {
		return CoverResp, err
	}

	// 参数转换
	for _, cover := range Cover {
		coverResp := respModel.CoverResp{
			//Todo:cover has no field nickname
			//Nickname: cover.Nickname,
			Avatar:   cover.Avatar,
			PostTime: tools.DecodeTime(cover.CreatedAt),
		}
		CoverResp = append(CoverResp, coverResp)
	}
	return CoverResp, nil
}

// 歌曲播放页的跳转操作[童年]
func GetPlayerChild(Jump int, CoverId int) (respModel.PlayerChildResp, error) {
	MysqlDB := setting.MysqlConn()
	Cover := statements.Cover{}

	if Jump == 0 {
		err := MysqlDB.Where("id<? and module=2", CoverId).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerChildResp{}, errors.New("已经是第一首")
		}
	} else if Jump == 1 {
		err := MysqlDB.Where("id>? and module=2", CoverId).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerChildResp{}, errors.New("已经是最后一首")
		}
	} else {
		err := MysqlDB.Where("id=? and module=2", CoverId).First(&Cover).Error
		if err != nil {
			return respModel.PlayerChildResp{}, errors.New("数据库出错")
		}
	}

	return respModel.PlayerChildResp{
		SongName: Cover.SongName,
		File:     Cover.File,
		Icon:     Cover.Avatar,
		WorkName: "作品名(写死)",
		//Nickname:Cover.Nickname,
	}, nil
}

// 歌曲播放页的跳转操作[普通]
func GetPlayerNormal(Jump int, CoverId int) (respModel.PlayerNormalResp, error) {
	MysqlDB := setting.MysqlConn()
	Cover := statements.Cover{}

	if Jump == 0 {
		err := MysqlDB.Where("id<? and module=1", CoverId).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerNormalResp{}, errors.New("已经是第一首")
		}
	} else if Jump == 1 {
		err := MysqlDB.Where("id>? and module=1", CoverId).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerNormalResp{}, errors.New("已经是最后一首")
		}
	} else {
		err := MysqlDB.Where("id=? and module=1", CoverId).First(&Cover).Error
		if err != nil {
			return respModel.PlayerNormalResp{}, errors.New("数据库出错")
		}
	}

	return respModel.PlayerNormalResp{
		SongName: Cover.SongName,
		File:     Cover.File,
		Avatar:   Cover.Avatar,
		//Nickname:Cover.Nickname,
	}, nil
}
