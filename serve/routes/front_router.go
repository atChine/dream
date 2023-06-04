package routes

import (
	"dream/serve/config"
	"dream/serve/routes/middleware"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// FrontRouter 前台页面接口路由
func FrontRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	err := r.SetTrustedProxies([]string{"*"})
	if err != nil {
		return nil
	}

	r.Use(middleware.Cors()) // 跨域中间件
	// 基于 cookies 存储 session
	store := cookie.NewStore([]byte(config.Cfg.Session.Salt))

	store.Options(sessions.Options{MaxAge: config.Cfg.Session.MaxAge})
	r.Use(sessions.Sessions(config.Cfg.Session.Name, store)) // Session 中间件 (使用 Redis 存储引擎)
	// 开发模式同时把日志写到控制台
	if config.Cfg.Server.AppMode == "debug" {
		r.Use(gin.Logger()) // gin 默认日志挺好看的
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	//无需监权的接口
	base := r.Group("/api/front")
	{
		article := base.Group("/article")
		{
			article.GET("/list", fArticleAPI.GetFrontList) // 前台文章列表
		}
	}
	return r
}
