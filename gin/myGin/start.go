package myGin

import "C"
import (
	"chatroom/postType"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lxn/win"
	webview "github.com/webview/webview_go"
)

const (
	ICON_SMALL = 0 // 对应 Windows SDK 的 ICON_SMALL
	ICON_BIG   = 1 // 对应 Windows SDK 的 ICON_BIG
)

var closeServer_main = false

func Start(port string, re redis.Conn, channel chan postType.PostRequest, end chan bool) *gin.Engine {

	gi := gin.Default()
	staticAbsoluteAddr, _ := filepath.Abs(path.Join("./dist"))
	gi.Static("/", staticAbsoluteAddr)
	fmt.Println("已挂载以下页面：", staticAbsoluteAddr)
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
			fmt.Println("软件后端已关闭")
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

		lpszName, _ := syscall.UTF16PtrFromString("./dist/favicon_256.ico") // 图标路径
		// 加载图标文件
		hIcon := win.LoadImage(
			0,
			lpszName,
			win.IMAGE_ICON,
			0, 0,
			win.LR_LOADFROMFILE|win.LR_DEFAULTSIZE,
		)
		if hIcon == 0 {
			log.Printf("无法加载图标: %v", win.GetLastError())
		} else {
			// 设置窗口图标
			win.SendMessage(hWnd, win.WM_SETICON, ICON_SMALL, uintptr(hIcon))
			win.SendMessage(hWnd, win.WM_SETICON, ICON_BIG, uintptr(hIcon))
		}
		defer func() {
			if hIcon != 0 {
				win.DestroyIcon(win.HICON(hIcon))
			}
		}()

		// win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 置顶
		win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE) // 不置顶
		style := win.GetWindowLong(hWnd, win.GWL_STYLE)                                       // 普通窗口样式
		// style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION      // 无边框
		style |= win.WS_SIZEBOX | win.WS_CAPTION                                  // 有边框
		win.SetWindowLong(hWnd, win.GWL_STYLE, style)                             // 设置样式
		win.MoveWindow(hWnd, int32(10), int32(10), int32(720), int32(1280), true) // 移动窗口位置和大小
		win.SendMessage(hWnd, win.WM_KEYDOWN, 0x0000007A, 0x20380001)             // 按下按键
		win.SendMessage(hWnd, win.WM_KEYUP, 0x0000007A, 0x003C0001)               // 抬起按键
		// 监听窗口消息
		// go func() {
		var msg win.MSG
		for {
			if win.GetMessage(&msg, 0, 0, 0) <= 0 {
				break
			}
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		}
		// }()
	}()

	defer w.Destroy()
	w.SetTitle("望子成龙聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	w.Navigate(addr)
	fmt.Println("软件页面已启动:", addr)
	// w.Navigate("http://localhost:12306/")
	w.Run()
}
