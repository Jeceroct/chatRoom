package config

import (
	"chatroom/error"

	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	RedisPassword string
	RedisAddr     string
	RedisDB       string
	GinPort       string

	ListenerLastLen    int64
	NumOfConcurrentMsg int

	RoomName string
}

var configPath = "./chatRoom.conf.json"

var payload Config

var content, err1 = os.ReadFile(configPath)
var err2 = json.Unmarshal(content, &payload)

func init() {
	for i := 0; ; i++ {
		if i > 4 {
			error.Exit("配置文件创建失败，请检查程序是否有写入权限", 1)
		}
		if err1 != nil || err2 != nil {
			os.WriteFile(configPath, []byte(`{
				"RedisPassword": "",
				"RedisAddr": "",
				"RedisDB": "0",
  			"GinPort": ":12306",
  			"ListenerLastLen": 0,
  			"NumOfConcurrentMsg": 10,
				"RoomName": ""
			}`), 0644)
			content, err1 = os.ReadFile(configPath)
			err2 = json.Unmarshal(content, &payload)
		} else {
			break
		}
	}

	// 设置默认值
	if payload.RedisPassword == "" {
		fmt.Println("Redis密码未设置(\"RedisPassword\")")
	}
	if payload.RedisAddr == "" {
		// error.Exit("Redis地址和端口是必需的，请检查配置文件", 1)
		fmt.Println("Redis地址未设置(\"RedisAddr\")")
	}
	if payload.RedisDB == "" {
		fmt.Println("Redis数据库未设置,将使用默认数据库(\"RedisDB\")")
	}
	if payload.GinPort == "" {
		fmt.Println("Gin端口未设置,将使用默认端口\":12306\"(\"GinPort\")")
		payload.GinPort = ":12306"
	}
	if payload.ListenerLastLen == 0 {
		fmt.Println("最后一条读取的消息位置未设置,将使用默认值0(\"ListenerLastLen\")")
		payload.ListenerLastLen = 0
	}
	if payload.NumOfConcurrentMsg == 0 {
		fmt.Println("消息同步加载数量未设置,将使用默认值10(\"NumOfConcurrentMsg\")")
		payload.NumOfConcurrentMsg = 10
	}
	if payload.RoomName == "" {
		fmt.Println("聊天室名称获取失败(\"RoomName\")")
	}
	fmt.Println("读取配置文件成功")
}

func RedisPassword() string {
	return payload.RedisPassword
}

func RedisAddr() string {
	return payload.RedisAddr
}

func RedisDB() string {
	return payload.RedisDB
}

func GinPort() string {
	return payload.GinPort
}

func ListenerLastLen() int64 {
	return payload.ListenerLastLen
}

func NumOfConcurrentMsg() int {
	return payload.NumOfConcurrentMsg
}

func RoomName() string {
	return payload.RoomName
}
