package postType

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var basePath = "./dist"

func TypeCheck(msg PostRequest, c *gin.Context) PostRequest {
	switch msg.Type {
	case "error":
		c.JSON(501, gin.H{
			"message": "不支持的消息类型",
		})
	case "image":
		msg.Context = handleImage(msg, c)
	case "file":
		// msg.Context = handleFile(msg, c)
	case "text":
	default:
		c.JSON(501, gin.H{
			"message": "不支持的消息类型",
		})
	}
	return msg
}

func HandleFile(msg PostRequest, c *gin.Context, re redis.Conn) string {
	// fileMsg := ParseFileContext(msg)
	path := "/download/" + msg.Context
	// 如果download文件夹不存在，则创建
	if _, err := os.Stat(basePath + "/download"); os.IsNotExist(err) {
		os.Mkdir(basePath+"/download", os.ModePerm)
	}
	// 如果路径下已有同名文件，则直接返回文件路径
	if _, err := os.Stat(basePath + path); err == nil {
		// 直接使用cmd打开文件
		go exec.Command("cmd", "/c", "start", basePath+path).Run()
		return path
	}
	// 如果路径下没有同名文件，则从redis中获取文件内容并保存到本地
	context, err := redis.String(re.Do("GET", msg.Context))
	if err != nil {
		fmt.Println("获取文件内容时发生错误：", err)
		c.JSON(501, gin.H{
			"message": "文件接收失败",
		})
	}
	fmt.Println("文件接收成功")
	file, err := os.Create(basePath + path)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		c.JSON(501, gin.H{
			"message": "创建文件失败",
		})
	}
	defer file.Close()

	data := []byte(context)
	if _, err := file.Write(data); err != nil {
		fmt.Println("保存文件时发生错误：", err)
	}
	go exec.Command("cmd", "/c", "start", basePath+path).Run()
	return path
}

func handleImage(msg PostRequest, c *gin.Context) string {
	// 提取Base64编码部分
	base64Data := msg.Context
	if strings.HasPrefix(base64Data, "base64,") {
		base64Data = base64Data[strings.Index(base64Data, ","):]
	}

	// 解码Base64
	data, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		fmt.Println("解析base64时发生错误：", err)
		fmt.Println("错误的base64字符串：", msg.Context[:50])
	}
	name := time.Now().Format("2004-03-22_05-04-05")
	// 如果temp文件夹不存在，则创建
	if _, err := os.Stat(basePath + "/temp"); os.IsNotExist(err) {
		os.Mkdir(basePath+"/temp", os.ModePerm)
	}
	path := "/temp/" + name
	// 检查路径下是否已有同名文件，若有则在文件名后添加数字
	i := 1
	for {
		if _, err := os.Stat(basePath + path); os.IsNotExist(err) {
			break
		}
		path = fmt.Sprintf("%s(%d)", path, i)
		i++
	}
	file, err := os.Create(basePath + path)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		c.JSON(501, gin.H{
			"message": "图片接收失败",
		})
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		fmt.Println("保存图片时发生错误：", err)
		c.JSON(501, gin.H{
			"message": "图片接收失败",
		})
	}
	return path
}
