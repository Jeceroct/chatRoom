package myGin

import (
	"chatroom/config"
	"chatroom/data"
	"chatroom/myUser"
	"chatroom/postType"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SettingPost(gi *gin.Engine) {
	// 更新用户信息
	gi.POST("/updateUserInfo", func(c *gin.Context) {
		var req myUser.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(905, gin.H{
				"msg": "参数错误",
			})
		}
		myUser.UpdateAvatar(req.Avatar)
		myUser.UpdateId(req.Id)
		myUser.UpdateName(req.Name)
		myUser.UpdateTitle(req.Title)
		myUser.UpdateTitleColor(req.TitleColor)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "更新用户信息成功",
		})
		// myUser.UploadUserInfo(req, c)
	})

	// 导入聊天记录
	gi.POST("/importData", func(c *gin.Context) {
		var jsonData []postType.PostRequest
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(905, gin.H{
				"msg": "聊天记录格式错误",
			})
			fmt.Println("聊天记录格式错误:", err)
			return
		}
		// 保存文件
		data.ImportData(jsonData)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "保存文件成功",
		})
	})

	// 导出聊天记录
	gi.POST("/exportData", func(c *gin.Context) {
		// 将data.json发回
		c.File("./data.json")
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "导出文件成功",
		})
	})

	// 退出登录
	gi.POST("/logout", func(c *gin.Context) {
		// 转到user_page状态
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "退出登录成功",
		})
		CloseServer_start <- true
		Page <- RoutePage.USER_PAGE
		// 删除user.conf.json文件
		myUser.DeleteUserInfo()
	})

	// 退出聊天室
	gi.POST("/leaveRoom", func(c *gin.Context) {
		// 转到address_page状态
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "退出聊天室成功",
		})
		CloseServer_start <- true
		Page <- RoutePage.ADDRESS_PAGE
		// 删除user.conf.json文件
		myUser.DeleteUserInfo()
		// 删除chatRoom.conf.json文件
		config.DeleteConfig()
		// 删除data.json文件
		data.DeleteData()
	})
}
