package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//和数据库相关的放到dao

// DB 全局 DB，所有地方都能用
var DB *gorm.DB

func InitDB() {
	dsn := "root:051128@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	// 赋值给全局 DB
	DB = db
	// 自动迁移表

}
