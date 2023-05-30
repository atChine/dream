package main

import (
	"love_blog/model"
	"love_blog/routes"
)

func main() {
	// 连接数据库
	model.InitDb()

	routes.InitRouter()
}
