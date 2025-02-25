package main

import (
	"chatroom/config"
	"chatroom/myGin"
	"chatroom/myRedis"
	"chatroom/myUser"
	"chatroom/postType"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var openAddress string

var isInitialized = make(chan bool, 1)

func main() {
	openAddress = "http://localhost" + config.GinPort() + "/"

	go func() {
		// 检查聊天室配置
		checkRoomAddr()
		for {
			if <-isInitialized {
				break
			}
			time.Sleep(1 * time.Second)
		}
		// 检查用户信息
		checkUserInfo()
		for {
			if <-isInitialized {
				break
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Println("所有配置加载成功")

		var reSend redis.Conn
		var reGet redis.Conn
		// 连接至远程数据库
		for {
			reSend = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
			reGet = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
			if reSend != nil && reGet != nil {
				break
			} else {
				fmt.Println("连接失败，正在重试...")
				time.Sleep(5 * time.Second)
				continue
			}
		}
		// 监听远程数据库消息列表
		channel := make(chan postType.PostRequest, config.NumOfConcurrentMsg())
		myRedis.StartListen(reGet, "chat", channel)
		// 启动本地服务
		myGin.Start(config.GinPort(), reSend, channel)
	}()

	// 循环保证程序不退出
	// for {
	// 	time.Sleep(1 * time.Second)
	// }

	myGin.StartView(openAddress)
}

func checkRoomAddr() {
	if config.RedisAddr() == "" || config.RedisPassword() == "" || config.RedisDB() == "" {
		fmt.Println("聊天室地址未设置")
		go myGin.ConnectRedis(config.GinPort(), isInitialized)
		// isInitialized = true
		return
	}

	isInitialized <- true
}

func checkUserInfo() {
	if myUser.GetInfo().Id == "" {
		fmt.Println("用户信息未设置")
		openAddress = "http://localhost" + config.GinPort() + "/login"
		reCheck := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
		go myGin.BeforeStart(config.GinPort(), isInitialized, reCheck)
		// isInitialized = true
		return
	}
	isInitialized <- true
}
