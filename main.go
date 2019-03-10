package main

import (
	"github.com/symoo/sf55-api/model"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/symoo/sf55-api/config"
	"github.com/symoo/sf55-api/router"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path")
)

func main() {
	// 初始化配置
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 初始化框架
	gin.SetMode(viper.getString("runmode"))
	r := gin.New()
	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()
    // 加载路由
	middlerwares := []gin.HandlerFunc{}
	router.Load(r, middlerwares...,)
	
	r.Run()
}


