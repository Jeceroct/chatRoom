package myGin

import "C"
import (
	"chatroom/postType"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lxn/win"
	webview "github.com/webview/webview_go"
)

var closeServer_main = false

func Start(port string, re redis.Conn, channel chan postType.PostRequest, end chan bool) *gin.Engine {

	gi := gin.Default()
	gi.Static("/", "./dist")
	Post(gi, re, channel)
	fmt.Println("软件后端已启动")
	server := &http.Server{
		Addr:    port,
		Handler: gi,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	for {
		if closeServer_main {
			server.Close()
			end <- true
			fmt.Println("聊天室连接成功")
			break
		}
	}
	return gi
}

func StartView(addr string) {
	debug := true
	w := webview.New(debug)

	hwnd := w.Window()
	go func() {
		hWnd := win.HWND(hwnd)
		// win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 置顶
		win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 不置顶
		style := win.GetWindowLong(hWnd, win.GWL_STYLE)                                       // 普通窗口样式
		// style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION      // 无边框
		style |= win.WS_SIZEBOX | win.WS_CAPTION                                  // 有边框
		win.SetWindowLong(hWnd, win.GWL_STYLE, style)                             // 设置样式
		win.MoveWindow(hWnd, int32(10), int32(10), int32(720), int32(1280), true) // 移动窗口位置和大小
		win.SendMessage(hWnd, win.WM_KEYDOWN, 0x0000007A, 0x20380001)             // 按下按键
		win.SendMessage(hWnd, win.WM_KEYUP, 0x0000007A, 0x003C0001)               // 抬起按键
	}()

	defer w.Destroy()
	w.SetTitle("望子成龙聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	fmt.Println("软件页面已启动:", addr)
	w.Navigate(addr)
	// w.Navigate("http://localhost:12306/")
	w.Run()
}
