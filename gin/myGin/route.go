package myGin

import (
	"chatroom/config"
	"fmt"
)

var Page = make(chan int, 5)

type routePage struct {
	STOP_PROCESS int
	ROOM_PAGE    int
	ADDRESS_PAGE int
	USER_PAGE    int
}

var RoutePage = &routePage{
	STOP_PROCESS: 0,
	ROOM_PAGE:    1,
	ADDRESS_PAGE: 2,
	USER_PAGE:    3,
}

func Route() {
	t1, t2 := -1, -1
	for {
		t1 = t2
		t2 = <-Page
		if t1 == t2 {
			continue
		}
		fmt.Println("收到页面切换请求: ", t2)
		switch t2 {
		case RoutePage.ROOM_PAGE:
			fmt.Println("进入聊天页面")
			Start(config.GinPort())
		case RoutePage.ADDRESS_PAGE:
			fmt.Println("进入连接页面")
			ConnectRedis(config.GinPort())
		case RoutePage.USER_PAGE:
			fmt.Println("进入用户设置页面")
			BeforeStart(config.GinPort())
		case RoutePage.STOP_PROCESS:
			fmt.Println("退出程序")
			return
		}
	}
}
