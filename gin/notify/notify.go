package notify

// import (
// 	"fmt"
// 	"syscall"

// 	"github.com/lxn/win"
// )

// // Windows API 常量与接口定义
// const (
// 	CLSID_ToastNotificationManager      = "{50AC103F-D235-4598-BBEF-98FE4D1A3AD4}"
// 	IID_INotificationActivationCallback = "{53E31837-6600-4A81-9395-75CFFE746F94}"
// )

// var (
// 	ole32                = syscall.NewLazyDLL("ole32.dll")
// 	procCoInitializeEx   = ole32.NewProc("CoInitializeEx")
// 	procCoCreateInstance = ole32.NewProc("CoCreateInstance")
// 	procCoUninitialize   = ole32.NewProc("CoUninitialize")
// )

// // ToastNotification 结构体封装通知逻辑
// type ToastNotification struct {
// 	appID string
// 	icon  string
// }

// func NewToastNotification(appID, iconPath string) *ToastNotification {
// 	return &ToastNotification{
// 		appID: appID,
// 		icon:  iconPath,
// 	}
// }

// // 初始化 COM 组件
// func (t *ToastNotification) initializeCOM() error {
// 	hr, _, _ := procCoInitializeEx.Call(
// 		0,
// 		win.COINIT_APARTMENTTHREADED,
// 	)
// 	if hr != win.S_OK {
// 		return fmt.Errorf("COM 初始化失败: 0x%X", hr)
// 	}
// 	return nil
// }

// // 发送通知
// func (t *ToastNotification) Show(title, message string) error {
// 	// 生成 Toast XML 模板
// 	xml := t.generateToastXML(title, message)

// 	// 调用 Windows API 发送通知
// 	// （此处需实现 XML 提交逻辑，以下为伪代码）
// 	hr := t.sendToastXML(xml)
// 	if hr != win.S_OK {
// 		return fmt.Errorf("发送通知失败: 0x%X", hr)
// 	}
// 	return nil
// }

// // 生成符合 Windows Toast 规范的 XML
// func (t *ToastNotification) generateToastXML(title, message string) string {
// 	return fmt.Sprintf(`
// <toast>
//     <visual>
//         <binding template="ToastGeneric">
//             <text>%s</text>
//             <text>%s</text>
//             <image src="file:///%s" placement="appLogoOverride"/>
//         </binding>
//     </visual>
// </toast>
//     `, title, message, t.icon)
// }

// // 释放 COM 资源
// func (t *ToastNotification) uninitialize() {
// 	procCoUninitialize.Call()
// }

// func main() {
// 	toast := NewToastNotification(
// 		"YourCompany.Chatroom",   // 唯一 AppID
// 		"./dist/favicon_256.ico", // 绝对路径
// 	)

// 	if err := toast.initializeCOM(); err != nil {
// 		panic(err)
// 	}
// 	defer toast.uninitialize()

// 	if err := toast.Show("新消息", "用户A: 你好！"); err != nil {
// 		panic(err)
// 	}
// }
