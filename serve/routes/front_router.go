package routes

import (
	"dream/serve/config"
	"dream/serve/routes/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FrontRouter 前台页面接口路由
func FrontRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	//err := r.SetTrustedProxies([]string{"*"})
	//if err != nil {
	//	return nil
	//}

	r.Use(middleware.Cors()) // 跨域中间件
	// 基于 cookies 存储 session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt))

	store.Options(sessions.Options{MaxAge: config.Cfg.Session.MaxAge})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store))
	//Session 中间件 (使用 Redis 存储引擎)
	//开发模式同时把日志写到控制台
	if config.Cfg.Server.AppMode == "debug" {
		r.Use(gin.Logger()) // gin 默认日志挺好看的
	}

	//无需监权的接口
	base := r.Group("/api/front")
	{
		article := base.Group("/article")
		{
			article.GET("/list", fArticleAPI.GetFrontList)      // 前台文章列表
			article.GET("/:id", fArticleAPI.GetFrontInfo)       //根据id查询文章详情
			article.GET("/archive", fArticleAPI.GetArchiveList) // 前台文章归档
			article.GET("/search", fArticleAPI.Search)          // 前台文章搜索
		}
	}
	return r
}
