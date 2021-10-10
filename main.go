package main

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/controller/auth"
	"git.100steps.top/100steps/healing2021_be/cron"
	"git.100steps.top/100steps/healing2021_be/models"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"git.100steps.top/100steps/healing2021_be/router"
	"github.com/fvbock/endless"
	"io/ioutil"
	"log"
	"syscall"
)

// @Title healing2021
// @Version 1.0
// @Description 2021治愈系

func main() {
	models.TableInit()
	routers:=router.SetupRouter()
	auth.Login()
	defer setting.DB.Close()
	defer setting.RedisClient.Close()
	var port string
	if tools.IsDebug() {
		//controller.LoadTestData()

		port = ":8003"
	} else {
		port = ":8001"
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
