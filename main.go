package main

import (
	"ChatRoom/config"
	"fmt"
	"log"
)

func main() {
	fmt.Println("!-------------!")
	fmt.Println("!----启动中----!")
	fmt.Println("!-----原木-----!")
	fmt.Println("!---倾情-制作---!")
	fmt.Println("!-------------!")

	// 初始化配置和Redis
	if err := config.Init(); err != nil {
		log.Fatalf("初始化失败: %v", err)
	}

	defer config.GetRedis().Close()

	StartView()
}
