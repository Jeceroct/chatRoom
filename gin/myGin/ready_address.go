package myGin

import (
	"chatroom/config"
	"chatroom/myRedis"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Addr     string `json:"address"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

var closeServer_address = false

func ConnectRedis(port string, end chan bool) {
	gi := gin.Default()
	gi.Static("/", "./dist")

	server := &http.Server{
		Addr:    port,
		Handler: gi,
	}

	addr_post(gi)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	for {
		if closeServer_address {
			server.Close()
			end <- true
			fmt.Println("聊天室连接成功")
			break
		}
	}
}

func addr_post(gi *gin.Engine) {
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
				closeServer_address = true
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
}
