package front

import (
	"dream/serve/utils/r"
	"github.com/gin-gonic/gin"
)

type Link struct {
}

// GetFrontList 获取前台友链链接
func (*Link) GetFrontList(c *gin.Context) {
	r.SuccessData(c, linkService.GetFrontList())
}
