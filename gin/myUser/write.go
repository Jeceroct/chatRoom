package myUser

import (
	"encoding/json"
	"os"
)

func UpdateId(id string) {
	payload.Id = id
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateName(name string) {
	payload.Name = name
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateLevel(level int) {
	payload.Level = level
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateAvatar(avatar string) {
	payload.Avatar = avatar
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateTitle(title string) {
	payload.Title = title
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdateTitleColor(titleColor string) {
	payload.TitleColor = titleColor
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func UpdatePhone(phone string) {
	payload.Phone = phone
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}

func DeleteUserInfo() {
	payload.Id = ""
	payload.Name = ""
	payload.Level = 0
	payload.Avatar = ""
	payload.Title = ""
	payload.TitleColor = ""
	payload.Phone = ""
	jsonData, _ := json.Marshal(payload)
	os.WriteFile(configPath, jsonData, 0644)
}
