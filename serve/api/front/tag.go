package front

import (
	"dream/serve/utils/r"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

// GetFrontList 前台标签列表
func (*Tag) GetFrontList(c *gin.Context) {
	r.SuccessData(c, tagService.GetFrontList())
}
