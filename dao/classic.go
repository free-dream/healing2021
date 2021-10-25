package dao

import (
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

// 获取前十的翻唱歌曲信息
func GetTop10() ([]respModel.ClassicResp, error) {
	MysqlDB := setting.MysqlConn()
	var ClassicResp []respModel.ClassicResp
	var ClassicList []statements.Classic

	err := MysqlDB.Order("likes desc").Limit(10).Find(&ClassicList).Error
	if err != nil {
		return ClassicResp, err
	}

	//格式转换
	for _, classic := range ClassicList {
		classicResp := respModel.ClassicResp{
			Name:  classic.SongName,
			Icon:  classic.Icon,
			Click: classic.Click,
		}
		ClassicResp = append(ClassicResp, classicResp)
	}
	return ClassicResp, nil
}

// 所有翻唱歌曲的信息
func GetLIst() ([]respModel.ClassicListResp, error) {
	MysqlDB := setting.MysqlConn()
	var ClassicListResp []respModel.ClassicListResp
	var ClassicList []statements.Classic

	err := MysqlDB.Find(&ClassicList).Error
	if err != nil {
		return ClassicListResp, err
	}

	for _, classic := range ClassicList {
		classicList := respModel.ClassicListResp{
			Name:   classic.SongName,
			Avatar: classic.Icon,
			Time:   classic.CreatedAt.String(),
		}
		ClassicListResp = append(ClassicListResp, classicList)
	}
	return ClassicListResp, nil
}

// 获取原唱的信息
func GetOriginInfo(Name string) (respModel.OriginInfoResp, error) {
	MysqlDB := setting.MysqlConn()
	Origin := statements.Classic{}

	err := MysqlDB.Where("song name=?", Name).First(&Origin).Error
	if err != nil {
		return respModel.OriginInfoResp{}, err
	}

	// 格式转换
	OriginInfoResp := respModel.OriginInfoResp{
		SongId: int(Origin.ID),
		Name:   Origin.SongName,
		Singer: Origin.Singer,
		Icon:   Origin.Icon,
	}
	return OriginInfoResp, nil
}
