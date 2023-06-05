package front

import (
	"dream/serve/model/req"
	"dream/serve/utils"
	"dream/serve/utils/r"
	"github.com/gin-gonic/gin"
	"log"
)

type Article struct{}

// GetFrontList 获取前台文章列表
func (*Article) GetFrontList(c *gin.Context) {
	log.Println("Request URL:", c.Request.URL.Path)
	r.SuccessData(c, articleService.GetFrontList(utils.BindQuery[req.GetFrontArts](c)))
}
