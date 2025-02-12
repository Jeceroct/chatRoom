package data

import (
	"chatroom/error"
	"chatroom/postType"
	"encoding/json"
	"fmt"
	"os"
)

var datas []postType.PostRequest

var configPath = "./data.json"

var content, err1 = os.ReadFile(configPath)
var err2 = json.Unmarshal(content, &datas)

func init() {
	for i := 0; ; i++ {
		if i > 4 {
			error.Exit("历史数据读取失败", 1)
		}
		if err1 != nil || err2 != nil {
			os.WriteFile(configPath, []byte("[]"), 0644)
			content, err1 = os.ReadFile(configPath)
			err2 = json.Unmarshal(content, &datas)
		} else {
			break
		}
	}
	fmt.Println("历史数据读取成功")
}

func GetData() []postType.PostRequest {
	return datas
}
