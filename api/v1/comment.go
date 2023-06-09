package v1

import (
	"dream/model"
	"dream/utils"
	"dream/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetCommentListFront 展示页面显示评论列表
func GetCommentListFront(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, pageNum = utils.HandlePageSizeAndPageNum(pageSize, pageNum)
	data, total, code := model.GetCommentListFront(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}

// GetCommentCount 获取评论数量
func GetCommentCount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total, code := model.GetCommentCount(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
