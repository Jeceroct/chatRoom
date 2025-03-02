// main.go
package main

import (
	"ChatRoom/config"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

// 用户信息结构体
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	TitleColor string `json:"titleColor"`
}

// 引用消息结构体
type Quote struct {
	From    User   `json:"from"`
	Type    string `json:"type"`
	Context string `json:"context"`
	Time    string `json:"time"`
}

// 消息结构体
type Message struct {
	Type    string `json:"type"`    // text or image
	Context string `json:"context"` // Text content or Image URL
	Time    string `json:"time"`    // Timestamp
	From    User   `json:"from"`    // Sender
	Quote   *Quote `json:"quote"`   // Quoted message (optional)
}

// **格式化普通消息**
func (m *Message) FormatMessage() string {
	return fmt.Sprintf("[%s][%s] %s: %s",
		m.Time, m.From.Title, m.From.Name, m.Context)
}

// **格式化带引用的消息**
func (m *Message) FormatWithQuote() string {
	if m.Quote != nil && m.Quote.Context != "" {
		return fmt.Sprintf("[%s] %s 回复 %s:\n> %s\n%s",
			m.Time, m.From.Name, m.Quote.From.Name, m.Quote.Context, m.Context)
	}
	return m.FormatMessage()
}

// **格式化图片消息**
func (m *Message) FormatImageMessage() string {
	return fmt.Sprintf("[%s][%s] %s 发送了一张图片: %s",
		m.Time, m.From.Title, m.From.Name, m.Context)
}

//// Redis 客户端
//var ctx = context.Background()
//var client = redis.NewClient(&redis.Options{
//	Addr: "localhost:6379",
//})

// 获取Redis客户端实例

// **加载历史消息**
func loadChatHistory(chatBox *widget.Entry) {
	redisClient := config.GetRedis()
	messages, err := redisClient.GetHistory()

	if err != nil {
		log.Println("加载历史消息失败:", err)
		return
	}

	var chatHistory string
	for _, msgStr := range messages {
		var msg Message
		err := json.Unmarshal([]byte(msgStr), &msg)
		if err != nil {
			log.Println("消息解析失败:", err)
			continue
		}

		// 选择格式化方式
		var formattedMessage string
		if msg.Type == "image" {
			//formattedMessage = msg.FormatImageMessage()
		} else if msg.Quote != nil {
			formattedMessage = msg.FormatWithQuote()
		} else {
			formattedMessage = msg.FormatMessage()
		}

		// 按行拼接消息
		chatHistory += formattedMessage + "\n"
	}

	// 设置聊天框内容
	chatBox.SetText(chatHistory)
}

func StartView() {

	// 创建 Fyne GUI 应用
	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme()) // 使用暗色主题

	w := a.NewWindow("Go 聊天室")

	// 修改窗口尺寸
	w.Resize(fyne.NewSize(600, 500)) // 调整为更大的窗口尺寸

	// 添加顶部标题
	title := widget.NewLabel("群聊：Go 学习小组")
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	//title.Theme().Size(18)

	// 聊天框（用于显示消息）
	chatBox := widget.NewMultiLineEntry()
	chatBox.Wrapping = fyne.TextWrapWord // 启用自动换行
	chatBox.SetMinRowsVisible(20)        // 显示更多行
	//chatBox.Disable()
	//chatBox.TextStyle = fyne.TextStyle{Bold: true} // 加粗字体
	//chatBox.TextColor = color.White                // 设置白色字体
	//chatBox.CursorColor = color.White              // 光标颜色
	//chatBox.Theme().Color(theme.ColorBlue, theme.VariantDark)

	// 创建带滚动条的聊天框容器
	chatScroll := container.NewScroll(chatBox)
	chatScroll.SetMinSize(fyne.NewSize(580, 400)) // 固定滚动区域大小

	// 加载历史消息
	loadChatHistory(chatBox)
	chatScroll.ScrollToBottom() // 添加自动滚动到底部

	// 输入框
	input := widget.NewEntry()
	// 输入框样式修改
	input.SetPlaceHolder("输入消息...")
	input.TextStyle = fyne.TextStyle{Italic: true} // 斜体提示文字
	input.Validator = nil                          // 移除默认验证
	input.MultiLine = true                         // 允许多行输入
	input.SetMinRowsVisible(2)                     // 输入框默认显示两行
	//input.TextColor = color.White  // 输入文字白色

	// 创建表情按钮（左边新增）
	emojiBtn := widget.NewButtonWithIcon("", theme.MenuIcon(), func() {
		// 表情选择功能（待实现）
	})
	emojiBtn.Importance = widget.MediumImportance

	// 发送按钮
	sendBtn := widget.NewButtonWithIcon("发送", theme.MailSendIcon(), func() {
		msgText := input.Text
		if msgText != "" {
			// 构造消息结构体
			msg := Message{
				Type:    "text",
				Context: msgText,
				Time:    time.Now().Format("1.2 15:04"),
				From: User{
					ID:         "000",
					Name:       "迷糊老师",
					Level:      10,
					Avatar:     "avatar.png",
					Title:      "高贵的群主",
					TitleColor: "#f5e50dbc",
				},
			}

			// 获取Redis客户端实例
			redisClient := config.GetRedis()

			// 发布消息
			// 序列化 JSON 并存入 Redis
			msgJSON, _ := json.Marshal(msg)

			if err := redisClient.PublishMessage(config.GetConfig().RoomKey, msgJSON); err != nil {
				log.Printf("消息发布失败: %v", err)
			}
			if err := redisClient.SaveMessage(msgJSON); err != nil {
				log.Printf("消息保存失败: %v", err)
			}
			//client.Publish(ctx, "chatroom", msgJSON)
			//client.RPush(ctx, "chat_history", msgJSON)

			input.SetText("")
		}
	})
	sendBtn.Importance = widget.HighImportance // 高亮显示按钮
	sendBtn.SetIcon(theme.MailSendIcon())      // 添加发送图标

	// 创建底部输入容器
	inputContainer := container.NewBorder(
		nil, nil,
		container.NewPadded(emojiBtn), // 左边表情按钮
		container.NewPadded(sendBtn),  // 右边发送按钮
		input,                         // 中间输入框
	)

	// 订阅 Redis 频道
	//pubsub := client.Subscribe(ctx, "chatroom")

	go func() {
		// 获取Redis客户端实例
		redisClient := config.GetRedis()
		ch := redisClient.Subscribe().Channel()
		//ch := pubsub.Channel()
		for msg := range ch {
			// 解析 JSON 消息
			var receivedMsg Message
			err := json.Unmarshal([]byte(msg.Payload), &receivedMsg)
			if err != nil {
				log.Println("消息解析失败:", err)
				continue
			}

			// 选择格式化方式
			var formattedMessage string
			if receivedMsg.Type == "image" {
				formattedMessage = receivedMsg.FormatImageMessage()
			} else if receivedMsg.Quote != nil {
				formattedMessage = receivedMsg.FormatWithQuote()
			} else {
				formattedMessage = receivedMsg.FormatMessage()
			}

			// **确保 UI 更新在主线程执行**
			//fyne.CurrentApp().Driver().CallOnMain(func() {
			chatBox.SetText(chatBox.Text + "\n" + formattedMessage)
			//})

			// 在消息更新后添加自动滚动逻辑（在订阅消息的 goroutine 中）
			//chatBox.SetText(chatBox.Text + "\n" + formattedMessage)
			chatScroll.ScrollToBottom() // 添加自动滚动到底部
		}
	}()

	//// 修改整体布局
	//w.SetContent(container.NewBorder(
	//	chatBox,        // 顶部
	//	inputContainer, // 底部
	//	input,          // 左边
	//	sendBtn,        // 右边
	//	chatScroll,     // 中间区域
	//))
	// 修改整体布局
	w.SetContent(container.NewBorder(
		container.NewCenter(title),          // 顶部
		container.NewPadded(inputContainer), // 底部
		nil,                                 // 左边
		nil,                                 // 右边
		chatScroll,                          // 中间区域
	))
	w.Resize(fyne.NewSize(500, 400))
	w.ShowAndRun()
}
