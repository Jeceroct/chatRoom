package myGin

import "C"
import (
	"chatroom/config"
	"chatroom/myRedis"
	"chatroom/myUser"
	"chatroom/postType"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/lxn/win"
	"github.com/polevpn/webview"
)

type Address struct {
	Addr     string `json:"address"`
	Password string `json:"password"`
	DB       string `json:"db"`
}
type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Title    string `json:"title"`
}

func Start(port string, re redis.Conn, channel chan postType.PostRequest) *gin.Engine {
	gi := gin.Default()
	gi.Static("/", "./dist")
	Post(gi, re, channel)
	fmt.Println("软件后端已启动")
	go gi.Run(port)
	return gi
}

func ConnectRedis(port string, end chan bool) {
	isInitialized := false
	gi := gin.Default()
	gi.Static("/", "./dist")
	gi.POST("/address", func(c *gin.Context) {
		var addr Address
		c.BindJSON(&addr)
		fmt.Println("连接请求: ", addr.Addr, addr.Password, addr.DB)
		reCheck := myRedis.Connect(addr.Addr, addr.Password, addr.DB, 0)
		if reCheck != nil {
			c.JSON(200, gin.H{
				"code": "200",
			})
			config.UpdateRedisAddr(addr.Addr)
			config.UpdateRedisPassword(addr.Password)
			config.UpdateRedisDB(addr.DB)
			reCheck.Close()
			go func() {
				time.Sleep(1 * time.Second)
				isInitialized = true
			}()
		} else {
			c.JSON(200, gin.H{
				"code": "501",
			})
		}
	})
	gi.POST("/getStatus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "checkRoom",
		})
	})

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
		if isInitialized {
			server.Close()
			end <- true
			fmt.Println("聊天室连接成功")
			break
		}
	}
}

func BeforeStart(port string, end chan bool, reCheck redis.Conn) {
	isInitialized := false
	gi := gin.Default()
	gi.Static("/", "./dist")

	gi.POST("/login", func(c *gin.Context) {
		var user1 User
		var user2 User
		c.BindJSON(&user1)
		fmt.Println("登录请求: ", user1.Id, user1.Password)
		context, err := reCheck.Do("GET", user1.Id)
		if err != nil {
			c.JSON(200, gin.H{
				"code": "501",
			})
			fmt.Println("未找到用户：", err)
		} else {
			c.JSON(200, gin.H{
				"code": "200",
			})
			jsonStr, _ := context.([]byte)
			json.Unmarshal(jsonStr, &user2)
			fmt.Println("查询到的用户：", user2)
			if user1.Password != user2.Password {
				c.JSON(200, gin.H{
					"code": "501",
				})
				return
			}
			var user myUser.User
			json.Unmarshal(jsonStr, &user)
			myUser.UpdateId(user.Id)
			myUser.UpdateName(user.Name)
			myUser.UpdateTitle(user.Title)
			myUser.UpdateLevel(user.Level)
			myUser.UpdatePhone(user.Phone)
			myUser.UpdateTitleColor(user.TitleColor)
			myUser.UpdateAvatar(user.Avatar)
			go func() {
				time.Sleep(1 * time.Second)
				isInitialized = true
			}()
		}
	})

	gi.POST("/signup", func(c *gin.Context) {
		var userP User
		c.BindJSON(&userP)
		fmt.Println("注册请求: ", userP.Id, userP.Password)
		ok, _ := redis.Bool(reCheck.Do("EXISTS", userP.Id))
		if ok {
			c.JSON(200, gin.H{
				"code": "501",
				"msg":  "用户已存在",
			})
		}
		con, _ := json.Marshal(userP)
		reCheck.Do("SET", userP.Id, con)
		reCheck.Do("SADD", "userList", userP.Id)
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "注册成功",
		})
	})

	gi.POST("/checkIdUsed", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		fmt.Println("检查请求: ", user.Id)
		ok, _ := redis.Bool(reCheck.Do("EXISTS", user.Id))
		if ok {
			c.JSON(200, gin.H{
				"code": "501",
				"msg":  "用户已存在",
			})
		} else {
			c.JSON(200, gin.H{
				"code": "200",
			})
		}
	})

	gi.POST("/getStatus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "checkUser",
		})
	})

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
		if isInitialized {
			server.Close()
			end <- true
			fmt.Println("聊天室连接成功")
			break
		}
	}
}

func StartView(addr string) {
	hide, debug := false, true
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
	w.SetTitle("望子成龙聊天室")
	// w.SetSize(720, 1280, webview.HintNone)
	fmt.Println("软件页面已启动:", addr)
	w.Navigate(addr)
	// w.Navigate("http://localhost:12306/")
	w.Run()
}
