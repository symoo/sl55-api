package router

import (
	"github.com/gin-gonic/gin"
	"github.com/symoo/sf55-api/handler/v1/user"
)

// Load 加载路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(mw...)

	v1 := g.Group("v1/")
	{
		v1.GET("login", user.Login)
	}

	return g
}