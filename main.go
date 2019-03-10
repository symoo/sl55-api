package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	r.Run()
}

func login(code string) string {
	r := code2Session(code)
	
	return ""
}

func code2Session(code string) interface{} {
	appid := ""
	secret := ""
	grantType := "authorization_code"
	return 1
}

func findByOpenID(openID string) {
	
}