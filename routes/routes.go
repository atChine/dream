package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	v1 "love_blog/api/v1"
	"love_blog/middleware"
	"love_blog/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	//r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUserList)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)
		//修改密码
		auth.PUT("admin/changepw/:id", v1.ChangePassword)
		// 分类模块的路由接口
		auth.GET("admin/category", v1.GetCateList)
		auth.POST("category/add", v1.AddCate)
		auth.PUT("category/:id", v1.EditCateById)
		auth.DELETE("category/:id", v1.DelCateById)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfoById)
		auth.GET("admin/article", v1.GetArtInfo)
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.EdiArtById)
		auth.DELETE("article/:id", v1.DelArtById)
		// 上传文件
		//auth.POST("upload", v1.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", v1.GetUserInfoById)
		auth.PUT("profile/:id", v1.EditUser)
		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteCommentById)
		//auth.PUT("checkcomment/:id", v1.CheckComment)
		//auth.PUT("uncheckcomment/:id", v1.UncheckComment)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfoById)
		router.GET("users", v1.GetUserList)

		// 文章分类信息模块
		router.GET("category", v1.GetCateList)
		router.GET("category/:id", v1.GetCateInfoById)

		// 文章模块
		router.GET("article", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetArtInfoByCate)
		router.GET("article/info/:id", v1.GetArtInfoById)

		// 登录控制模块
		//router.POST("login", v1.Login)
		//router.POST("loginfront", v1.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", v1.GetUserInfoById)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetCommentById)
		router.GET("commentfront/:id", v1.GetArtCommentList)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
