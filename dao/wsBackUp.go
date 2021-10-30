package dao

import (
    "encoding/json"
    "git.100steps.top/100steps/healing2021_be/models/statements"
    "git.100steps.top/100steps/healing2021_be/pkg/setting"
    "git.100steps.top/100steps/healing2021_be/pkg/respModel"
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

func GetAllSysMsg(uid uint) ([]respModel.SysMsg, error){
    db := setting.MysqlConn()
    resp := make([]respModel.SysMsg,1)
    err := db.Model(&statements.Sysmsg{}).Where("uid = ? and is_send = 1", uid).Order("created_at desc").Find(&resp).Error
    return resp, err
}

func GetAllUsrMsg(uid uint) ([]respModel.UsrMsg, error){
    db := setting.MysqlConn()
    resp := make([]respModel.UsrMsg,1)
    err := db.Model(&statements.Usrmsg{}).Where("uid = ? and is_send = 1", uid).Order("created_at desc").Find(&resp).Error
    return resp, err
}
