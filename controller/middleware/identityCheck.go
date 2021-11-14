package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

func IdentityCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rUrl := ctx.Request.URL.Path
		session := sessions.Default(ctx)
		openid := session.Get("openid")
		if startWith(rUrl, "/auth") || startWith(rUrl, "/wx") {
			ctx.Next()
			return
		}
		if openid == nil {
			if startWith(rUrl, "/api") {
				ctx.JSON(401, e.ErrMsgResponse{Message: "fail to authenticate"})
				ctx.Abort()
				return
			} else {
				redirect := ctx.Query("redirect")
				var url string
				if tools.IsDebug() {
					url = "https://healing2021.test.100steps.top/wx/jump2wechat?redirect=" + redirect
				} else {
					url = "https://healing2021.100steps.top/wx/jump2wechat?redirect=" + redirect
				}
				ctx.Redirect(302, url)
				ctx.Abort()
				return
			}
		}
		//微信登陆任务初始化
		userid := tools.GetUserid(ctx)
		taskTableInit(userid)
		//
		ctx.Next()
	}
}

func startWith(rUrl string, uri string) bool {
	if len(uri) > len(rUrl) {
		return false
	}
	rUrl = rUrl[0:len(uri)]
	return rUrl == uri
}

//任务表初始化
func taskTableInit(userid int) {
	normaltasks := []int{1, 2, 3}
	dao.GenerateTasktable(normaltasks, userid)
}
