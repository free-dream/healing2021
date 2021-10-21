package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

// 获取所有翻唱信息（全表查找，后面优化
func GetCoverList(Name string)  ([]respModel.CoverResp, error){
	MysqlDB := setting.MysqlConn()
	var CoverResp []respModel.CoverResp
	var Cover []statements.Cover
	
	err := MysqlDB.Where("").Find(&Cover).Error
	if err != nil {
		return CoverResp, err
	}
	
	// 参数转换
	for _, cover := range CoverResp {
		coverResp := respModel.CoverResp{
			Nickname: cover.Nickname,
			Avatar: cover.Avatar,
			PostTime: cover.PostTime,
		}
		CoverResp = append(CoverResp, coverResp)
	}
	return CoverResp, nil
}

// 歌曲播放页的跳转操作[童年]
func GetPlayerChild(Jump int, CoverId int)  (respModel.PlayerChildResp,error){
	return respModel.PlayerChildResp{}, nil
}

// 歌曲播放页的跳转操作[普通]
func GetPlayerNormal(Jump int, CoverId int)  (respModel.PlayerNormalResp,error){
	return respModel.PlayerNormalResp{}, nil
}