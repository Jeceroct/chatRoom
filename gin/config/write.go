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
