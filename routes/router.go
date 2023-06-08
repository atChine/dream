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
		//文章模块
		router.GET("article", v1.GetArt)                // 获取全部文章列表 / 搜索title模糊查询
		router.GET("article/list/:id", v1.GetArtByCate) // 按照cate查询文章
		router.GET("article/info/:id", v1.GetInfoById)  // GetInfoById 查询单个文章信息
		//用户信息
		router.GET("user/:id", v1.GetUserById) // 根据id获取详细信息
	}

	_ = r.Run(utils.HttpPort)
}
