package dao

import (
	"errors"
	"fmt"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

// 获取当前歌曲的信息
func GetPlayerInfo(CoverId int) (respModel.PlayerResp, error) {
	MysqlDB := setting.MysqlConn()
	Cover := statements.Cover{}
	err := MysqlDB.Where("id=?", CoverId).First(&Cover).Error
	if err != nil {
		fmt.Println(err)
		return respModel.PlayerResp{}, err
	}


	// 找到作品名
	Classic := statements.Classic{}
	err = MysqlDB.Where("song_name=?", Cover.SongName).First(&Classic).Error
	if err != nil {
		fmt.Println(err)
		return respModel.PlayerResp{}, err
	}

	// 参数转换
	return respModel.PlayerResp{
		CoverId:  int(Cover.ID),
		Nickname: Cover.Nickname,
		Icon:     Cover.Avatar,
		File:     Cover.File,
		Name:     Cover.SongName,
		WorkName: Classic.WorkName,
	}, nil
}

// 歌曲播放页的跳转操作[童年]
func GetPlayerChild(Jump int, CoverId int) (respModel.PlayerResp, error) {
	MysqlDB := setting.MysqlConn()
	Cover := statements.Cover{}

	if Jump == 0 {
		err := MysqlDB.Where("id<? and module=?", CoverId, 2).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerResp{}, errors.New("已经是第一首")
		}
	} else if Jump == 1 {
		err := MysqlDB.Where("id>? and module=?", CoverId, 2).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerResp{}, errors.New("已经是最后一首")
		}
	} else {
		return respModel.PlayerResp{}, errors.New("jump参数出错")
	}

	// 找到作品名
	Classic := statements.Classic{}
	err := MysqlDB.Where("song_ame=?", Cover.SongName).First(&Classic).Error
	if err != nil {
		return respModel.PlayerResp{}, err
	}

	return respModel.PlayerResp{
		CoverId:  int(Cover.ID),
		File:     Cover.File,
		Name:     Cover.SongName,
		Nickname: Cover.Nickname,
		Icon:     Cover.Avatar,
		WorkName: Classic.WorkName,
	}, nil
}

// 歌曲播放页的跳转操作[普通]
func GetPlayerNormal(Jump int, CoverId int) (respModel.PlayerResp, error) {
	MysqlDB := setting.MysqlConn()
	Cover := statements.Cover{}

	if Jump == 0 {
		err := MysqlDB.Where("id<? and module=?", CoverId, 1).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerResp{}, errors.New("已经是第一首")
		}
	} else if Jump == 1 {
		err := MysqlDB.Where("id>? and module=?", CoverId, 1).Order("id desc").First(&Cover).Error
		if err != nil {
			return respModel.PlayerResp{}, errors.New("已经是最后一首")
		}
	} else {
		return respModel.PlayerResp{}, errors.New("jump参数出错")
	}

	return respModel.PlayerResp{
		CoverId:  int(Cover.ID),
		File:     Cover.File,
		Name:     Cover.SongName,
		Nickname: Cover.Nickname,
		Icon:     Cover.Avatar,
		WorkName: "",
	}, nil
}
