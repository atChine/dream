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

	//无需监权的接口
	//base := r.Group("/api/front")
	//{
	//	base.POST("/login", userAuthAPI.Login) // 登录
	//}
	return r
}
