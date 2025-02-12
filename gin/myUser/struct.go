package myUser

import (
	"chatroom/error"

	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Id         string
	Name       string
	Level      int
	Avatar     string
	Title      string
	TitleColor string
	Phone      string
}

var configPath = "./user.conf.json"

var payload User

var content, err1 = os.ReadFile(configPath)
var err2 = json.Unmarshal(content, &payload)

func init() {
	if err1 != nil || err2 != nil {
		error.Exit("读取用户配置失败，请检查用户配置文件", 1)
	}

	// 设置默认值
	if payload.Id == "" {
		error.Exit("用户Id是必需的，请检查用户配置(\"Id\")", 1)
	}
	if payload.Name == "" {
		error.Exit("用户名称是必需的，请检查用户配置(\"Name\")", 1)
	}
	if payload.Level < 0 {
		payload.Level = 0
	}
	if payload.Avatar == "" {
		fmt.Println("用户头像未设置(\"Avatar\")")
	}
	if payload.Title == "" {
		fmt.Println("用户头衔未设置(\"Title\")")
		payload.Title = "新用户"
	}
	if payload.TitleColor == "" {
		fmt.Println("用户头衔颜色未设置(\"TitleColor\")")
		payload.TitleColor = "#8d0df5bc"
	}
	if payload.Phone == "" {
		fmt.Println("用户手机号未设置(\"Phone\")")
	}

	fmt.Println("读取用户配置成功:", payload)
}

func Userfy(str string) User {
	str = strings.Replace(strings.Replace(str, "{", "", -1), "}", "", -1)
	m := strings.Split(str, " ")

	level, _ := strconv.Atoi(m[2])
	return User{
		Id:         m[0],
		Name:       m[1],
		Level:      level,
		Avatar:     m[3],
		Title:      m[4],
		TitleColor: m[5],
		Phone:      m[6],
	}
}

func GetInfo() User {
	return payload
}
