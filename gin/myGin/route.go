package myGin

import (
	"chatroom/config"
	"fmt"
)

type routePage struct {
	StopProcess  int
	StartPage    int
	ReadyAddress int
	ReadyUser    int
}

var RoutePage = &routePage{
	StopProcess:  0,
	StartPage:    1,
	ReadyAddress: 2,
	ReadyUser:    3,
}

func Route(page chan int) {
	t1, t2 := 0, 0
	for {
		t1 = t2
		t2 = <-page
		if t1 == t2 {
			continue
		}
		switch t2 {
		case RoutePage.StartPage:
			fmt.Println("进入聊天页面")
			Start(config.GinPort(), page)
		case RoutePage.ReadyAddress:
			fmt.Println("进入连接页面")
			ConnectRedis(config.GinPort(), page)
		case RoutePage.ReadyUser:
			fmt.Println("进入用户设置页面")
			BeforeStart(config.GinPort(), page)
		case RoutePage.StopProcess:
			return
		}
	}
}
