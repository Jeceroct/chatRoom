package myGin

import "C"
import (
	"chatroom/config"
	"chatroom/myRedis"
	"chatroom/postType"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/energye/systray"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lxn/win"
	webview "github.com/webview/webview_go"
)

const (
	ICON_SMALL = 0 // 对应 Windows SDK 的 ICON_SMALL
	ICON_BIG   = 1 // 对应 Windows SDK 的 ICON_BIG
)

var (
	debug        = true
	flashSignal  = make(chan bool, 1) // 闪烁控制信号
	stopSignal   = true               // 停止闪烁信号
	originalIcon []byte
	alertIcon    []byte
)

var CloseServer_start = make(chan bool, 1)

func Start(port string) *gin.Engine {

	gi := gin.Default()
	staticAbsoluteAddr, _ := filepath.Abs(path.Join("./dist"))
	gi.Static("/", staticAbsoluteAddr)
	fmt.Println("已挂载以下页面：", staticAbsoluteAddr)

	go func() {
		server = &http.Server{
			Addr:    port,
			Handler: gi,
		}
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		for {
			if <-CloseServer_start {
				server.Close()
				if config.RedisAddr() == "" {
					Page <- RoutePage.ADDRESS_PAGE
				} else {
					Page <- RoutePage.USER_PAGE
				}
				break
			}
		}
	}()

	var reSend redis.Conn
	var reGet redis.Conn
	// 连接至远程数据库
	reSend = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
	reGet = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
	if reSend == nil || reGet == nil {
		fmt.Println("连接远程数据库失败")
		CloseServer_start <- true
		return nil
	}

	// 监听远程数据库消息列表
	channel := make(chan postType.PostRequest, config.NumOfConcurrentMsg())
	StartListen(reGet, "chat", channel, CloseServer_start)

	Post(gi, reSend, channel, CloseServer_start)
	SettingPost(gi)

	return gi
}

var hWnd win.HWND
var hIcon win.HANDLE

func StartView(addr string) {
	w := webview.New(debug)

	hwnd := w.Window()
	go func() {
		hWnd = win.HWND(hwnd)

		lpszName, _ := syscall.UTF16PtrFromString("./dist/favicon_256.ico") // 图标路径
		// 加载图标文件
		hIcon = win.LoadImage(
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

		// 置顶
		// win.SetWindowPos(hWnd, win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE)

		// 不置顶
		win.SetWindowPos(hWnd, win.HWND_NOTOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE)

		// 普通窗口样式
		style := win.GetWindowLong(hWnd, win.GWL_STYLE)

		// 无边框
		// style &= ^win.WS_SIZEBOX & ^win.WS_CAPTION

		// 有边框
		style |= win.WS_SIZEBOX | win.WS_CAPTION

		// 去除最大化按钮
		style &^= win.WS_MAXIMIZEBOX

		// 设置样式
		win.SetWindowLong(hWnd, win.GWL_STYLE, style)

		// 刷新窗口
		win.MoveWindow(hWnd, int32(10), int32(10), int32(720), int32(1280), true)
		win.InvalidateRect(hWnd, nil, true)

		// 监听窗口消息
		originalWndProc := win.GetWindowLongPtr(hWnd, win.GWLP_WNDPROC)
		newWndProc := syscall.NewCallback(func(hWnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
			switch msg {
			// 阻止默认关闭行为
			case win.WM_CLOSE:
				win.ShowWindow(hWnd, win.SW_HIDE)
				fmt.Println("窗口已隐藏")
				return 0
			// 处理窗口激活消息
			case win.WM_ACTIVATE:
				if wParam == win.WA_ACTIVE || wParam == win.WA_CLICKACTIVE {
					stopSignal = true
					notifySignal = false
				} else {
					stopSignal = false
					notifySignal = true
				}
			}
			return win.CallWindowProc(originalWndProc, hWnd, msg, wParam, lParam)
		})
		win.SetWindowLongPtr(hWnd, win.GWLP_WNDPROC, newWndProc)
	}()

	defer w.Destroy()
	w.SetTitle("望子成龙聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	w.Navigate(addr)
	systray.Run(onReady, onExit)
	fmt.Println("软件页面已启动:", addr)
	// w.Navigate("http://localhost:12306/")
	w.Run()
}

// 托盘图标
func onReady() {
	iconPath, _ := filepath.Abs("./dist/favicon_256.ico")
	originalIcon, _ = os.ReadFile(iconPath)
	alertIconPath, _ := filepath.Abs("./dist/favicon_alert.ico")
	alertIcon, _ = os.ReadFile(alertIconPath)
	systray.SetIcon(originalIcon)
	systray.SetTitle("望子成龙小学")
	systray.SetTooltip("望子成龙小学")
	systray.SetOnClick(func(menu systray.IMenu) {
		// 点击托盘图标时执行的操作
		win.ShowWindow(hWnd, win.SW_SHOW)
		win.SetForegroundWindow(hWnd)
	})
	mQuit := systray.AddMenuItem("退出", "退出程序")
	mQuit.Click(func() {
		CloseServer_start <- true
		systray.Quit()
		os.Exit(0)
	})

	// 启动消息监听
	go func() {
		for range flashSignal {
			go startFlash()
		}
	}()
}

func onExit() {
	if hIcon != 0 {
		win.DestroyIcon(win.HICON(hIcon))
	}
	// 清理托盘图标
	systray.Quit()
}

func startFlash() {
	for {
		if stopSignal {
			return
		}
		systray.SetIcon(alertIcon)
		time.Sleep(500 * time.Millisecond)
		systray.SetIcon(originalIcon)
		time.Sleep(500 * time.Millisecond)
	}
}
