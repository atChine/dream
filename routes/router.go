package routes

import (
	"dream/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	router := r.Group("api/v1")
	{
		router.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hello",
			})
		})
	}

	_ = r.Run(utils.HttpPort)
}
