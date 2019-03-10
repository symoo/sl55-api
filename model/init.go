package model

import (
	"fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/spf13/viper"
)

// Database 数据库连接
type Database struct {
	Self *gorm.DB
	Docker *gorm.DB
}

var DB *Database

func open(user passwd, addr, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@TCP(%s)/%s?charset=utf8*parseTime=%t&loc=%s", user, passwd, addr, name, true, "Local")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

// Init 初始化数据库
func (db *Database) Init() {
	DB := &Database {
		Self: Get
	}
}