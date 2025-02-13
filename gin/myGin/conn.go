package myGin

import "C"
import (
	"chatroom/postType"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/polevpn/webview"
)

func Connect(port string, re redis.Conn, channel chan postType.PostRequest) *gin.Engine {
	gi := gin.Default()
	Post(gi, re, channel)
	gi.Static("/", "./dist")
	fmt.Println("软件后端已启动")
	go gi.Run(port)
	startView("http://localhost" + port)
	// gi.Run(port)
	return gi
}

func startView(addr string) {
	hide, debug := false, true
	w := webview.New(720, 1280, hide, debug)
	// hwnd := syscall.Handle(w.Window())
	// syscall.Syscall()

	defer w.Destroy()
	w.SetTitle("望子成龙小学聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	w.Navigate(addr)
	fmt.Println("软件页面已启动")
	w.Run()
}
