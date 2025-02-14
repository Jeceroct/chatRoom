package myGin

import "C"
import (
	"chatroom/postType"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lxn/win"
	"github.com/polevpn/webview"
)

func Start(port string, re redis.Conn, channel chan postType.PostRequest) *gin.Engine {
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
	hide, debug := false, false
	w := webview.New(0, 0, hide, debug)

	hwnd := w.Window()
	go func() {
		hWnd := win.HWND(hwnd)
		// win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 置顶
		win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 不置顶
		style := win.GetWindowLong(hWnd, win.GWL_STYLE)                                       // 普通窗口样式
		// style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION      // 无边框
		style |= win.WS_SIZEBOX | win.WS_CAPTION                                 // 有边框
		win.SetWindowLong(hWnd, win.GWL_STYLE, style)                            // 设置样式
		win.MoveWindow(hWnd, int32(10), int32(10), int32(450), int32(900), true) // 移动窗口位置和大小
		win.SendMessage(hWnd, win.WM_KEYDOWN, 0x0000007A, 0x20380001)            // 按下按键
		win.SendMessage(hWnd, win.WM_KEYUP, 0x0000007A, 0x003C0001)              // 抬起按键
	}()

	defer w.Destroy()
	w.SetTitle("望子成龙小学聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	w.Navigate(addr)
	fmt.Println("软件页面已启动")
	w.Run()
}
