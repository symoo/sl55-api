package user

import (
	"github.com/gin-gonic/gin"
)

// Login 使用小程序传递的 code 换取 session key
func Login(c *gin.Context) {
	c.String(200, "login")
}