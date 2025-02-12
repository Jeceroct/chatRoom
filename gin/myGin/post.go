package myGin

import (
	"chatroom/config"
	"chatroom/data"
	"chatroom/myRedis"
	"chatroom/myUser"
	"chatroom/postType"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func Post(gi *gin.Engine, re redis.Conn, channel chan postType.PostRequest) {

	// 接收消息内容并发往redis
	gi.POST("/send", func(c *gin.Context) {
		var req postType.PostRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(401, gin.H{
				"message": "消息格式错误",
			})
			return
		}
		if req.Type == "file" {
			file := postType.ParseFileContext(req)
			var err error
			req.Context, err = uploadFile(re, file)
			if err != nil {
				fmt.Println("向redis写入文件时失败", err)
				c.JSON(500, gin.H{
					"message": "文件上传失败",
				})
			}
		}
		jsonMsg, _ := json.Marshal(req)
		_, err := re.Do("RPUSH", "chat", jsonMsg)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "消息发送失败,Redis连接错误",
			})
			for {
				re = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
				if re != nil {
					break
				}
				time.Sleep(5 * time.Second)
			}
			return
		}
		c.JSON(200, gin.H{
			"message": "消息已发送: " + req.Context,
		})
	})

	// 从redis获取消息
	gi.POST("/get", func(c *gin.Context) {
		var res []postType.PostRequest
		// 如果channel已满，则一次性取出所有消息再返回，否则取一条就返回
		if len(channel) == cap(channel) {
			for {
				if len(channel) == 0 {
					break
				}
				msg := <-channel
				msg = postType.TypeCheck(msg, c)
				res = append(res, msg)
			}
		} else {
			msg := <-channel
			msg = postType.TypeCheck(msg, c)
			res = append(res, msg)
		}
		fmt.Println("收到消息: ", res)
		data.UpdateData(res)
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	// 获取用户信息
	gi.POST("/getUserInfo", func(c *gin.Context) {
		c.JSON(200, myUser.GetInfo())
	})

	// 获取历史消息
	gi.POST("/getHistory", func(c *gin.Context) {
		c.JSON(200, data.GetData())
	})

	// 下载文件
	gi.POST("/download", func(c *gin.Context) {
		var msg postType.PostRequest
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(401, gin.H{
				"message": "消息格式错误",
			})
			return
		}
		path := postType.HandleFile(msg, c, re)
		c.JSON(200, path)
	})
}

func uploadFile(re redis.Conn, file postType.FileType) (string, error) {
	// 先检查redis中是否已经有同名文件
	reRead := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB())
	for i := 0; ; i++ {
		_, err := reRead.Do("GET", file.Title)
		if err == nil {
			break
		} else {
			file.Title = file.Title + "(" + fmt.Sprint(i) + ")"
		}
	}
	reRead.Close()

	// 提取Base64编码部分(base64,位置后的部分)
	base64Data := file.Context
	if strings.HasPrefix(base64Data, "base64,") {
		base64Data = base64Data[strings.Index(base64Data, ","):]
	}

	// 解码Base64
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		fmt.Println("解码Base64时错误：", err)
		return "", err
	}
	_, err1 := re.Do("SET", file.Title, data)
	return file.Title, err1
}
