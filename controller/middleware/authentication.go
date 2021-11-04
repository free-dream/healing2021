package middleware

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() func(*gin.Context) {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		nickname := session.Get("nickname")
		isAdministrator := dao.Authentication(nickname.(string))
		if isAdministrator {
			ctx.Next()
		} else {
			ctx.JSON(403, "无管理员权限")
			ctx.Abort()
		}

	}
}
