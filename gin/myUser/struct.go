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
	Id         string `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	TitleColor string `json:"titleColor"`
	Phone      string `json:"phone"`
}

var configPath = "./user.conf.json"

var payload User

var content, err1 = os.ReadFile(configPath)
var err2 = json.Unmarshal(content, &payload)

func init() {
	for i := 0; ; i++ {
		if i > 4 {
			error.Exit("用户配置文件创建失败，请检查程序是否有写入权限", 1)
		}
		if err1 != nil || err2 != nil {
			os.WriteFile(configPath, []byte(`{
				"Level": 0,
				"Avatar": "",
				"Title": "新用户",
				"TitleColor": "#8d0df5bc",
				"Phone": ""
			}`), 0644)
			content, err1 = os.ReadFile(configPath)
			err2 = json.Unmarshal(content, &payload)
		} else {
			break
		}
	}

	// 设置默认值
	if payload.Id == "" {
		// error.Exit("用户Id是必需的，请检查用户配置(\"Id\")", 1)
		fmt.Println("用户未登录(\"Id\")")
	}
	if payload.Name == "" {
		// error.Exit("用户名称是必需的，请检查用户配置(\"Name\")", 1)
		fmt.Println("用户未登录(\"Name\")")
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
		UpdateTitle(payload.Title)
	}
	if payload.TitleColor == "" {
		fmt.Println("用户头衔颜色未设置(\"TitleColor\")")
		payload.TitleColor = "#8d0df5bc"
		UpdateTitleColor(payload.TitleColor)
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
