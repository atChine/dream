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
		router.GET("users", v1.GetUsers)       //  搜索用户
		router.POST("user/add", v1.AddUser)    // 增加用户
		// 文章分类信息模块
		router.GET("category", v1.GetCate)         //获取全部标签
		router.GET("category/:id", v1.GetCateInfo) // 查询分类信息
		// 个人信息
		router.GET("profile/:id", v1.GetProfileById) //获取个人设置信息
		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront) //前台登录
		// 评论模块
		router.GET("commentfront/:id", v1.GetCommentListFront) //获取评论列表
		router.GET("commentcount/:id", v1.GetCommentCount)     //获取评论数量
	}
	// 后台接口
	auth := r.Group("api/v1")
	{
		// 用户
		auth.GET("admin/users", v1.GetUsers)       //获取全部用户
		auth.PUT("user/:id", v1.EditUser)          // 设置用户信息
		auth.DELETE("user/:id", v1.DeleteUserById) //删除用户
		//重置用户密码
		auth.PUT("admin/restpw/:id", v1.ResetUserPassword) // 重置用户密码
		//分类模块
		auth.GET("admin/category", v1.GetCate)     //查询分类列表
		auth.POST("category/add", v1.AddCategory)  // 增加分类标签
		auth.PUT("category/:id", v1.EditCate)      // 编辑标签
		auth.DELETE("category/:id", v1.DeleteCate) //删除标签
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetInfoById) //获取查询单个文章信息
		auth.GET("admin/article", v1.GetArt)               // 获取文章列表
		auth.POST("article/add", v1.AddArticle)            //新增文章
		auth.PUT("article/:id", v1.EditArt)                //修改文章

	}
	_ = r.Run(utils.HttpPort)
}
