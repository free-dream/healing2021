package dao

import (
	state "git.100steps.top/100steps/healing2021_be/models/statements"
	resp "git.100steps.top/100steps/healing2021_be/pkg/respModel"
)

// 将数据库中的用户信息 User 进行提取转换为 UserInfo
func TransformUserInfo(OneUser state.User) resp.UserInfo {
	UserInfos := resp.UserInfo{
		Id:            int(OneUser.ID),
		Nackname:      OneUser.Nickname,
		Avatar:        OneUser.Avatar,
		AvatarVisible: OneUser.AvatarVisible,
	}

	return UserInfos
}
