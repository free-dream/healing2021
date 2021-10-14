package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"git.100steps.top/100steps/healing2021_be/pkg/e"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
)

func IdentityCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		rUrl := c.Request.URL.Path
		session := sessions.Default(c)
		openid := session.Get("openid")

		if startWith(rUrl, "/auth") || startWith(rUrl, "/wx") || startWith(rUrl, "/api/broadcast") {
			c.Next()
			return
		}
		if openid == nil {
			if startWith(rUrl, "/api") {
				c.JSON(401, e.ErrMsgResponse{Message: "fail to authenticate"})
				c.Abort()
				return
			} else {
				redirect := c.Query("redirect")
				var url string
				if tools.IsDebug() {
					url = "https://healing2020.100steps.top/test/wx/jump2wechat?redirect=" + redirect
				} else {
					url = "https://healing2020.100steps.top/wx/jump2wechat?redirect=" + redirect
				}
				c.Redirect(302, url)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func startWith(rUrl string, uri string) bool {
	if tools.IsDebug() {
		uri = "/test" + uri
	}
	if len(uri) > len(rUrl) {
		return false
	}
	rUrl = rUrl[0:len(uri)]
	return rUrl == uri
}
