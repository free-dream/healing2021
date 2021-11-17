package middleware

import (
	"git.100steps.top/100steps/healing2021_be/dao"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"github.com/gin-gonic/gin"
)

func TaskInit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//微信登陆任务初始化
		userid := tools.GetUserid(ctx)
		taskTableInit(userid)
		//
		ctx.Next()
	}
}

//任务表初始化
func taskTableInit(userid int) {
	normaltasks := []int{1, 2, 3}
	dao.GenerateTasktable(normaltasks, userid)
}
