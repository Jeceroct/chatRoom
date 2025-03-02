package myGin

import (
	"chatroom/config"
	"chatroom/myRedis"
	"chatroom/myUser"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Title    string `json:"title"`
}

var closeServer_user = false

func BeforeStart(port string, end chan bool) {
	gi := gin.Default()
	gi.Static("/", "./dist")

	user_post(gi)

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
		if closeServer_user {
			server.Close()
			end <- true
			fmt.Println("聊天室连接成功")
			break
		}
	}
}

func user_post(gi *gin.Engine) {
	gi.POST("/login", func(c *gin.Context) {
		var user1 User
		var user2 User
		c.BindJSON(&user1)
		fmt.Println("登录请求: ", user1.Id, user1.Password)
		reCheck := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
		context, err := reCheck.Do("GET", user1.Id)
		reCheck.Close()
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
				closeServer_user = true
			}()
		}
	})

	gi.POST("/signup", func(c *gin.Context) {
		var userP User
		c.BindJSON(&userP)
		fmt.Println("注册请求: ", userP.Id, userP.Password)
		reCheck := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
		ok, _ := redis.Bool(reCheck.Do("EXISTS", userP.Id))
		reCheck.Close()
		if ok {
			c.JSON(200, gin.H{
				"code": "501",
				"msg":  "用户已存在",
			})
		}
		con, _ := json.Marshal(userP)
		reCheck = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
		_, err1 := reCheck.Do("SET", userP.Id, con)
		_, err2 := reCheck.Do("SADD", "userList", userP.Id)
		reCheck.Close()
		if err1 != nil || err2 != nil {
			c.JSON(501, gin.H{
				"msg": "注册失败,可能是redis连接失效",
			})
		} else {
			c.JSON(200, gin.H{
				"code": "200",
				"msg":  "注册成功",
			})
		}
	})

	gi.POST("/checkIdUsed", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		fmt.Println("检查请求: ", user.Id)
		reIdCheck := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0)
		ok, _ := redis.Bool(reIdCheck.Do("EXISTS", user.Id))
		reIdCheck.Close()
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
}
