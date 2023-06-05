package front

import (
	"dream/serve/utils/r"
	"github.com/gin-gonic/gin"
)

type Category struct{}

// GetFrontList 前台分类列表
func (*Category) GetFrontList(c *gin.Context) {
	r.SuccessData(c, categoryService.GetFrontList())
}
