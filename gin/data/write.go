package data

import (
	"chatroom/config"
	"chatroom/postType"
	"encoding/json"
	"os"
)

func UpdateData(newData []postType.PostRequest) {
	datas = append(datas, newData...)
	jsonData, _ := json.Marshal(datas)
	os.WriteFile(configPath, jsonData, 0644)
}

func ImportData(newData []postType.PostRequest) {
	datas = newData
	jsonData, _ := json.Marshal(datas)
	os.WriteFile(configPath, jsonData, 0644)
	config.UpdateListenerLastLen(0)
}

func DeleteData() {
	os.Remove(configPath)
	datas = nil
}
