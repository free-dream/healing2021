package dao

import (
	state "git.100steps.top/100steps/healing2021_be/models/statements"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
)

// 将数据库中的用户信息 User 进行提取转换为 UserInfo(用于响应点歌)
func TransformUserInfo(OneUser state.User, selectionId int) resp.UserInfo {
	info := ""
	var err error
	if selectionId != -1 {
		info, err = GetSelectorInfo(selectionId)
		if err != nil {
			return resp.UserInfo{}
		}
	}

	UserInfos := resp.UserInfo{
		Id:            int(OneUser.ID),
		Nickname:      OneUser.Nickname,
		Avatar:        OneUser.Avatar,
		AvatarVisible: OneUser.AvatarVisible,
		Sex:           OneUser.Sex,
		Remark:        info,
	}

	return UserInfos
}
