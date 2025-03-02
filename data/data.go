// data/data.go
package data

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	mu    sync.Mutex
	items []Message
	path  = "./chat_data.json"
)

type Message struct {
	Content string `json:"content"`
	Sender  string `json:"sender"`
}

func Load() error {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil // 首次运行无文件正常
	}
	return json.Unmarshal(content, &items)
}

func Save(msg Message) error {
	mu.Lock()
	defer mu.Unlock()

	items = append(items, msg)
	content, _ := json.MarshalIndent(items, "", "  ")
	return os.WriteFile(path, content, 0644)
}

func History() []Message {
	return items
}

// data/data.go

//package data
//
//import (
//	"ChatRoom/posttype"
//	"sync"
//)
//
//var (
//	mu   sync.Mutex
//	msgs []*posttype.PostRequest
//)
//
//func SaveMessage(msg *posttype.PostRequest) {
//	mu.Lock()
//	defer mu.Unlock()
//	msgs = append(msgs, msg)
//}
//
//func History() []*posttype.PostRequest {
//	return msgs
//}
