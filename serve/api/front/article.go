package front

import (
	"dream/serve/model/req"
	"dream/serve/utils"
	"dream/serve/utils/r"
	"github.com/gin-gonic/gin"
)

type Article struct{}

// GetFrontList 获取前台文章列表
func (*Article) GetFrontList(c *gin.Context) {
	r.SuccessData(c, articleService.GetFrontList(utils.BindQuery[req.GetFrontArts](c)))
}

// GetFrontInfo 根据id查询文章详情
func (*Article) GetFrontInfo(c *gin.Context) {
	r.SuccessData(c, articleService.GetFrontInfo(c, utils.GetIntParam(c, "id")))
}

// Search 前台文章搜索
func (*Article) Search(c *gin.Context) {
	r.SuccessData(c, articleService.Search(utils.BindQuery[req.KeywordQuery](c)))
}
