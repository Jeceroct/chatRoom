package main

import (
	"chatroom/config"
	"chatroom/myGin"
	"chatroom/myUser"
	"fmt"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

var openAddress string

func main() {
	registerAppID()
	openAddress = "http://localhost" + config.GinPort() + "/"

	go func() {
		// 检查聊天室配置
		checkRoomAddr()
		// 检查用户信息
		checkUserInfo()
		// 启动本地服务
		myGin.Page <- myGin.RoutePage.ROOM_PAGE
		myGin.Route()
	}()

	myGin.StartView(openAddress)
}

func checkRoomAddr() {
	if config.RedisAddr() == "" || config.RedisDB() == "" {
		fmt.Println("聊天室地址未设置")
		myGin.Page <- myGin.RoutePage.ADDRESS_PAGE
	}
}

func checkUserInfo() {
	if myUser.GetInfo().Id == "" {
		fmt.Println("用户信息未设置")
		myGin.Page <- myGin.RoutePage.USER_PAGE
	}
}

func registerAppID() {
	keyPath := `Software\Classes\AppUserModelId\hrx.chatRoom`
	k, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	if err != nil {
		panic(err)
	}
	defer k.Close()
	k.SetStringValue("DisplayName", "望子成龙小学")
	iconAbsolutePath, _ := filepath.Abs("./dist/favicon_256.ico")
	k.SetStringValue("IconUri", iconAbsolutePath)
}
