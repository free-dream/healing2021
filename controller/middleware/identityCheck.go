package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

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

	}
}

func startWith(rUrl string, uri string) bool {
	if len(uri) > len(rUrl) {
		return false
	}
	rUrl = rUrl[0:len(uri)]
	return rUrl == uri
}
