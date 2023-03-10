package main

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/dao"
	"io/ioutil"
	"log"
	"syscall"

	"git.100steps.top/100steps/healing2021_be/controller"
	"git.100steps.top/100steps/healing2021_be/cron"
	"git.100steps.top/100steps/healing2021_be/models"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"git.100steps.top/100steps/healing2021_be/router"
	"github.com/fvbock/endless"
)

// @Title healing2021
// @Version 1.0
// @Description 2021治愈系

func main() {
	models.TableInit()
	if tools.IsDebug() {
		/*statements.TableClean()
		models.AddClassic()
		models.AddDevotion()*/
	}

	//启动点赞后台和点赞message后台
	go controller.LikeDaemon()
	go controller.MsgDaemon()

	// 动态推荐页的缓存更新
	go dao.UpdateMomentPage()

	routers := router.SetupRouter()

	defer setting.DB.Close()
	defer setting.TokenGetCli.Close()
	defer setting.RedisClient.Close()
	var port string
	if tools.IsDebug() {
		port = ":8008"
	} else {
		port = ":8005"
	}

	c := cron.CronInit()
	go c.Start()
	defer c.Stop()
	server := endless.NewServer(port, routers)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Actual pid is %d", pid)
		ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}
