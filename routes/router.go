package routes

import (
	v1 "dream/api/v1"
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
		router.GET("article", v1.GetArt) // 获取全部文章列表 / 搜索title模糊查询
	}

	_ = r.Run(utils.HttpPort)
}
