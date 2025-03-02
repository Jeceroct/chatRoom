// posttype/message.go
package posttype

import (
	"fmt"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	TitleColor string `json:"titleColor"`
	Phone      string `json:"phone"`
}

type Quote struct {
	From    User   `json:"from"`
	Type    string `json:"type"`
	Context string `json:"context"`
	Time    string `json:"time"`
}

type PostRequest struct {
	Type    string `json:"type"`
	Context string `json:"context"`
	Time    string `json:"time"`
	From    User   `json:"from"`
	Quote   Quote  `json:"quote"`
}

// 新增消息格式化方法
func (p *PostRequest) FormatMessage() string {
	return fmt.Sprintf("[%s][%s] %s: %s",
		p.Time,
		p.From.Title,
		p.From.Name,
		p.Context)
}

// 新增带引用的消息格式化方法
func (p *PostRequest) FormatWithQuote() string {
	if p.Quote.Context != "" {
		return fmt.Sprintf("[%s] %s 回复 %s:\n> %s\n%s",
			p.Time,
			p.From.Name,
			p.Quote.From.Name,
			p.Quote.Context,
			p.Context)
	}
	return p.FormatMessage()
}
