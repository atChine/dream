package v1

import (
	"dream/model"
	"dream/utils"
	"dream/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetCate 获取全部标签
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, pageNum = utils.HandlePageSizeAndPageNum(pageSize, pageNum)
	data, code, total := model.GetCate(pageSize, pageNum)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
