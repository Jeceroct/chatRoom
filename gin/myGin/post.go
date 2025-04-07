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
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-toast/toast"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/text/encoding/unicode"
)

// 设置通知间隔时间变量
var notifyPeriod = true
var notifySignal = false

func Post(gi *gin.Engine, re redis.Conn, channel chan postType.PostRequest, closeServer_start chan bool) {

	// 接收消息内容并发往redis
	gi.POST("/send", func(c *gin.Context) {
		var req postType.PostRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(901, gin.H{
				"message": "未知的消息格式",
			})
			return
		}

		// 检查发送信息的用户是否存在
		reget := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
		ok, _ := redis.Bool(reget.Do("EXISTS", req.From.Id))
		if !ok {
			fmt.Println("用户信息不存在！")
			c.JSON(903, gin.H{
				"err": "用户信息不存在，请先注册再发送消息",
			})
			reget.Close()
			return
		}
		// 提出头像
		var user User
		context, _ := reget.Do("GET", req.From.Id)
		json.Unmarshal(context.([]byte), &user)
		user.Avatar = req.From.Avatar
		jsonUser, _ := json.Marshal(user)
		_, err := reget.Do("SET", req.From.Id, jsonUser)
		if err != nil {
			fmt.Println("用户头像提取失败", err)
		}
		reget.Close()
		req.From.Avatar = ""

		if req.Type == "file" {
			file := postType.ParseFileContext(req)
			var err error
			req.Context, err = uploadFile(file)
			if err != nil {
				fmt.Println("向redis写入文件时失败", err)
				c.JSON(902, gin.H{
					"message": "文件上传失败",
				})
			}
		}
		jsonMsg, _ := json.Marshal(req)

		// 将消息写入redis
		_, err = re.Do("RPUSH", "chat", jsonMsg)
		if err != nil {
			fmt.Println(err)
			c.JSON(501, gin.H{
				"message": "消息发送失败,Redis连接失败",
			})
			re = myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
			if re == nil {
				closeServer_start <- true
				Page <- RoutePage.ADDRESS_PAGE
			}
			return
		}

		c.JSON(200, gin.H{
			"code":    200,
			"message": "消息已发送: " + req.Context,
		})
	})

	// 从redis获取消息
	gi.POST("/get", func(c *gin.Context) {
		var res []postType.PostRequest

		// 如果channel已满，则一次性取出所有消息再返回，否则取一条就返回
		if len(channel) == cap(channel) {
			for {
				msg := <-channel
				msg.From.Avatar = handleAvatar(msg)

				if isAbort(c) {
					fmt.Println("客户端已取消请求")
					select {
					case channel <- msg:
						fmt.Println("消息回放成功")
					default:
						fmt.Println("警告！channel已满，消息丢失")
					}
					return
				}
				msg = postType.TypeCheck(msg, c)
				res = append(res, msg)
			}
		} else {
			msg := <-channel
			msg.From.Avatar = handleAvatar(msg)

			if isAbort(c) {
				fmt.Println("客户端已取消请求")
				channel <- msg
				return
			}
			msg = postType.TypeCheck(msg, c)
			res = append(res, msg)
		}
		data.UpdateData(res)

		// 间隔5秒发送通知
		notify(res)
		fmt.Println("收到消息: ", toUTF8(res[0].From.Name), toUTF8(res[0].Context))

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

	// 打开图片
	gi.POST("/openImg", func(c *gin.Context) {
		var msg postType.PostRequest
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(901, gin.H{
				"message": "未知的消息格式",
			})
			return
		}
		postType.OpenImg(msg.Context)
		c.JSON(200, gin.H{
			"message": "图片已打开",
		})
	})

	// 下载文件
	gi.POST("/download", func(c *gin.Context) {
		var msg postType.PostRequest
		if err := c.BindJSON(&msg); err != nil {
			c.JSON(901, gin.H{
				"message": "未知的消息格式",
			})
			return
		}
		path := postType.HandleFile(msg, c)
		c.JSON(200, path)
	})

	gi.POST("/getStatus", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "start",
		})
	})
}

func uploadFile(file postType.FileType) (string, error) {
	// 删除file.title中的空格
	file.Title = strings.ReplaceAll(file.Title, " ", "")
	// 先检查redis中是否已经有同名文件
	reRead := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 1, 3)
	if reRead == nil {
		return "", fmt.Errorf("redis连接失败")
	}
	for i := 0; ; i++ {
		ok, _ := redis.Bool(reRead.Do("GET", file.Title))
		if ok {
			file.Title = file.Title + "(" + fmt.Sprint(i) + ")"
		} else {
			break
		}
	}
	reRead.Close()

	// 提取Base64编码部分(base64,位置后的部分)
	base64Data := file.Context[strings.Index(file.Context, ",")+1:]
	// 解码Base64
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		fmt.Println("解码Base64时错误：", err)
		return "", err
	}
	reSend := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 1, 3)
	_, err1 := reSend.Do("SET", file.Title, data)
	reSend.Close()
	return file.Title, err1
}

func notify(res []postType.PostRequest) {
	// 发送闪烁信号
	select {
	case flashSignal <- true:
	default:
	}

	if notifyPeriod && notifySignal {
		notifyPeriod = false
		notification := toast.Notification{
			AppID:   "hrx.chatRoom",
			Title:   toUTF8(res[0].From.Name),
			Message: toUTF8(res[0].Context),
		}
		if err := notification.Push(); err != nil {
			panic(err)
		}
		go func() {
			time.Sleep(3 * time.Second)
			notifyPeriod = true
		}()
	}
}

func toUTF8(s string) string {
	utf8Bytes, _ := unicode.UTF8.NewEncoder().Bytes([]byte(s))
	return string(utf8Bytes)
}

// TODO 判断客户端已取消请求
func isAbort(c *gin.Context) bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}

// 处理头像
func handleAvatar(msg postType.PostRequest) string {
	path := "./dist/avatar/" + msg.From.Id
	// 若不存在则创建文件夹并写入头像文件
	if _, err := os.Stat("./dist/avatar/"); os.IsNotExist(err) {
		os.Mkdir("./dist/avatar/", os.ModePerm)
	}
	if _, err := os.Stat(path); err != nil {
		// 提取Base64编码部分
		base64Data := msg.From.Avatar[strings.Index(msg.From.Avatar, ",")+1:]
		data, _ := base64.StdEncoding.DecodeString(base64Data)
		file, err := os.Create(path)
		file.Write(data)
		if err != nil {
			fmt.Println("写入头像文件失败", err)
			return msg.From.Avatar
		}
		file.Close()
	}
	return "/avatar/" + msg.From.Id
}
