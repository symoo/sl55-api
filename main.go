package main

import (
	"github.com/gin-gonic/gin"
	"github.com/symoo/sf55-api/router"
)

func main() {
	r := gin.New()
	
	middlerwares := []gin.HandlerFunc{}

	router.Load(r, middlerwares...,)
	
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})
	
	r.Run()
}


