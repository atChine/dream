package routes

import (
	"dream/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	// 前台接口
	router := r.Group("api/v1")
	{
		router.GET("/")
	}

	_ = r.Run(utils.HttpPort)
}
