package postType

import (
	"chatroom/config"
	"chatroom/myRedis"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/h2non/filetype"
)

var basePath = "./dist"

func TypeCheck(msg PostRequest, c *gin.Context) PostRequest {
	switch msg.Type {
	case "error":
		c.JSON(901, gin.H{
			"message": "不支持的消息类型",
		})
	case "image":
		msg.Context = handleImage(msg, c)
	case "file":
		// msg.Context = handleFile(msg, c)
	case "text":
	default:
		c.JSON(901, gin.H{
			"message": "不支持的消息类型",
		})
	}
	return msg
}

func HandleFile(msg PostRequest, c *gin.Context) string {
	// fileMsg := ParseFileContext(msg)
	path := "/download/" + msg.Context
	// 如果download文件夹不存在，则创建
	if _, err := os.Stat(basePath + "/download"); os.IsNotExist(err) {
		os.Mkdir(basePath+"/download", os.ModePerm)
	}
	// 如果路径下已有同名文件，则直接返回文件路径
	if _, err := os.Stat(basePath + path); err == nil {
		// 直接使用cmd打开文件
		absPath, _ := filepath.Abs(basePath + path)
		cmd := exec.Command("cmd", "/c", "start", absPath)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		go cmd.Run()
		return path
	}
	// 如果路径下没有同名文件，则从redis中获取文件内容并保存到本地
	reGet := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 1, 3)
	defer reGet.Close()
	context, err := redis.String(reGet.Do("GET", msg.Context))
	if err != nil {
		fmt.Println("获取文件内容时发生错误：", err)
		c.JSON(903, gin.H{
			"message": "文件接收失败",
		})
	}
	fmt.Println("文件接收成功")
	file, err := os.Create(basePath + path)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		c.JSON(904, gin.H{
			"message": "创建文件失败",
		})
	}
	defer file.Close()

	data := []byte(context)
	if _, err := file.Write(data); err != nil {
		fmt.Println("保存文件时发生错误：", err)
	}
	absPath, _ := filepath.Abs(basePath + path)
	cmd := exec.Command("cmd", "/c", "start", absPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	go cmd.Run()
	return path
}

func handleImage(msg PostRequest, c *gin.Context) string {
	// 提取Base64编码部分
	base64Data := msg.Context[strings.Index(msg.Context, ",")+1:]

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
		path = fmt.Sprintf("%s%d", path, i)
		i++
	}
	file, err := os.Create(basePath + path)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		c.JSON(903, gin.H{
			"message": "图片接收失败",
		})
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		fmt.Println("保存图片时发生错误：", err)
		c.JSON(903, gin.H{
			"message": "图片接收失败",
		})
	}
	return path
}

func OpenImg(path string) {
	img := strings.Replace(basePath, "./", "", -1) + strings.Replace(path, "/", "\\", -1)
	absolutePath, _ := filepath.Abs(img)
	// 检测文件类型
	file, err := os.Open(absolutePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件头部字节（261字节足够检测常见类型）
	buf := make([]byte, 261)
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	// 匹配文件类型
	kind, err := filetype.Match(buf)
	if err != nil || kind == filetype.Unknown {
		fmt.Println("无法识别文件类型")
		return
	}

	tempDir := os.TempDir()
	tempFile := filepath.Join(tempDir, "temp_image."+kind.Extension)
	os.Remove(tempFile)
	cmd := exec.Command("cmd", "/C", "copy", absolutePath, tempFile, "&&", "start", "", tempFile)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	go cmd.Run()
}
