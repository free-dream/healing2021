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

	err := MysqlDB.Order("click desc").Limit(10).Find(&ClassicList).Error
	if err != nil {
		return ClassicResp, err
	}

	//格式转换
	for _, classic := range ClassicList {
		classicResp := respModel.ClassicResp{
			ClassicId: int(classic.ID),
			Name:      classic.SongName,
			Icon:      classic.Icon,
			Click:     classic.Click,
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
			ClassicId: int(classic.ID),
			Name:      classic.SongName,
			Icon:      classic.Icon,
			WorkName:  classic.WorkName,
		}
		ClassicListResp = append(ClassicListResp, classicList)
	}
	return ClassicListResp, nil
}

// 获取原唱的信息
func GetOriginInfo(ClassicId int) (respModel.OriginInfoResp, error) {
	MysqlDB := setting.MysqlConn()
	Origin := statements.Classic{}

	err := MysqlDB.Where("id=?", ClassicId).First(&Origin).Error
	if err != nil {
		return respModel.OriginInfoResp{}, err
	}

	ClassURL, err := GetClassicUrlById(ClassicId)
	if err != nil {
		return respModel.OriginInfoResp{}, err
	}

	// 格式转换
	OriginInfoResp := respModel.OriginInfoResp{
		ClassicURL: ClassURL,
		SongName:   Origin.SongName,
		Singer:     Origin.Singer,
		Icon:       Origin.Icon,
		WorkName:   Origin.WorkName,
	}
	return OriginInfoResp, nil
}

// 通过歌名找 classic_id (要求给的童年歌曲歌名不能重复)
type ClassicId struct {
	Id int `gorm:"classic_id"`
}

func GetClassicIdByName(SongName string) (int, error) {
	db := setting.MysqlConn()
	ClassicId := ClassicId{}
	err := db.Model(&statements.Classic{}).Where("song_name=?", SongName).Scan(&ClassicId).Error
	return ClassicId.Id, err
}

// 通过 classic_id 找 file_url
type ClassicUrl struct {
	File string `gorm:"file"`
}

func GetClassicUrlById(ClassicId int) (string, error) {
	db := setting.MysqlConn()
	tmpClassicUrl := ClassicUrl{}
	err := db.Select("file").Model(&statements.Classic{}).Where("id=?", ClassicId).Scan(&tmpClassicUrl).Error
	return tmpClassicUrl.File, err
}

// 通过 classic_id 找 song_name
type ClassicName struct {
	SongName string `gorm:"song_name"`
}

func GetClassicSongNameById(ClassicId int) (string, error) {
	db := setting.MysqlConn()
	tmpClassicSongName := ClassicName{}
	err := db.Select("song_name").Model(&statements.Classic{}).Where("id=?", ClassicId).Scan(&tmpClassicSongName).Error
	return tmpClassicSongName.SongName, err
}
