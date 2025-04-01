package config

import (
	"encoding/json"
	"os"
)

func UpdateListenerLastLen(lastLen int64) {
	payload.ListenerLastLen = lastLen
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateRedisAddr(addr string) {
	payload.RedisAddr = addr
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateRedisPassword(password string) {
	payload.RedisPassword = password
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateRedisDB(db string) {
	payload.RedisDB = db
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateGinPort(port string) {
	payload.GinPort = port
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateNumOfConcurrentMsg(num int) {
	payload.NumOfConcurrentMsg = num
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateRoomName(name string) {
	payload.RoomName = name
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func DeleteConfig() {
	os.Remove(configPath)
	payload = Config{}
}
