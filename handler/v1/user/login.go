package user

import (
	"github.com/gin-gonic/gin"
)



// Login 使用小程序传递的 code 换取 session key
func Login(c *gin.Context) {
	c.Bind()
	code := c.PostForm("code")
	if code == "" {
		c.JSON(200, gin.H{
			"code": 10001,
			"msg": "need code",
			"data": struct{},
		})
	}

	c.JSON(200, gin.H{
		"access_token": "1",
		"token_type": "bearer",
		"expire_in": 36000,
	})
}