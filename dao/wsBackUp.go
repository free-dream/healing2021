package dao

import (
	"encoding/json"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/respModel"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
)

func SysBackUp(msg respModel.SysMsg, isSend int) error {
	msgClone := statements.Sysmsg{}
	msgStr, _ := json.Marshal(msg)
	json.Unmarshal(msgStr, &msgClone)
	db := setting.MysqlConn()
	msgClone.IsSend = isSend
	return db.Model(&statements.Sysmsg{}).Create(&msgClone).Error
}

func UsrBackUp(msg respModel.UsrMsg, isSend int) error {
	msgClone := statements.Usrmsg{}
	msgStr, _ := json.Marshal(msg)
	json.Unmarshal(msgStr, &msgClone)
	db := setting.MysqlConn()
	msgClone.IsSend = isSend
	return db.Model(&statements.Usrmsg{}).Create(&msgClone).Error
}

func GetAllSysMsg(uid uint) ([]respModel.Sysmsg, error) {
	db := setting.MysqlConn()
	resp := make([]respModel.Sysmsg, 1)
	err := db.Model(&statements.Sysmsg{}).Where("uid = ?", uid).Order("created_at desc").Find(&resp).Error
	return resp, err
}

func GetAllUsrMsg(uid uint) ([]respModel.Usrmsg, error) {
	db := setting.MysqlConn()
	resp := make([]respModel.Usrmsg, 1)
	err := db.Model(&statements.Usrmsg{}).Where("from_user = ?", uid).Order("created_at desc").Find(&resp).Error
	return resp, err
}

func SysUpdate(uid uint) error {
    db := setting.MysqlConn()
    err := db.Model(&statements.Sysmsg{}).Where("uid = ?", uid).Update("is_send", 1).Error
    return err
}

func UsrUpdate(uid uint) error {
    db := setting.MysqlConn()
    err := db.Model(&statements.Usrmsg{}).Where("from_user = ?", uid).Update("is_send", 1).Error
    return err
}
