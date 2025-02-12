package data

import (
	"chatroom/postType"
	"encoding/json"
	"os"
)

func UpdateData(newData []postType.PostRequest) {
	datas = append(datas, newData...)
	jsonData, _ := json.Marshal(datas)
	os.WriteFile(configPath, jsonData, 0644)
}
