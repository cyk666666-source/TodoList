package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
)

func main() {

	dao.InitDB()
	err := dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	//链接数据库
	r := routers.SetupRouter()
	_ = r.Run()
}
