package main

import (
	"chatroom/config"
	"chatroom/myGin"
	"chatroom/myRedis"
	"chatroom/postType"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	var reSend redis.Conn
	var reGet redis.Conn
	// 连接至远程数据库
	for {
		reSend = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
		reGet = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
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
	myGin.Connect(config.GinPort(), reSend, channel)
}
