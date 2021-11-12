package main

import (
	"fmt"
	"git.100steps.top/100steps/healing2021_be/cron"
	"git.100steps.top/100steps/healing2021_be/models"
	"git.100steps.top/100steps/healing2021_be/models/statements"
	"git.100steps.top/100steps/healing2021_be/pkg/setting"
	"git.100steps.top/100steps/healing2021_be/pkg/tools"
	"git.100steps.top/100steps/healing2021_be/router"
	"git.100steps.top/100steps/healing2021_be/sandwich"
	"github.com/fvbock/endless"
	"io/ioutil"
	"log"
	"strconv"
	"syscall"
	"time"
)

// @Title healing2021
// @Version 1.0
// @Description 2021治愈系

func main() {
	models.TableInit()
	if tools.IsDebug() {
		statements.TableClean()
		time.Sleep(time.Second * 2)
		models.FakeData()
	}
	models.AddClassic()
	models.AddFakeHomeC()
	models.AddFakeHomeS()

	routers := router.SetupRouter()

	defer setting.DB.Close()
	defer setting.RedisClient.Close()
	var port string
	if tools.IsDebug() {
		//controller.LoadTestData()
		for i := 0; i < 10; i++ {
			sandwich.PutInHotSong(tools.EncodeSong(
				tools.HotSong{
					SongName: "歌曲" + strconv.Itoa(i),
					Language: "中文",
					Style:    "轻松",
				}))
			sandwich.PutInStates("状态" + strconv.Itoa(i))
			sandwich.PutInSearchWord("热词" + strconv.Itoa(i))
		}

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
