package myUser

import (
	"chatroom/config"
	"chatroom/myRedis"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func UploadUserInfo(userInfo User, c *gin.Context) {
	reSet := myRedis.Connect(config.RedisAddr(), config.RedisPassword(), config.RedisDB(), 0, 3)
	userJson, _ := json.Marshal(userInfo)
	if _, err := reSet.Do("SET", userInfo.Id, userJson); err != nil {
		c.JSON(905, gin.H{
			"msg": "上传用户信息失败",
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传用户信息成功",
	})
	reSet.Close()
}
