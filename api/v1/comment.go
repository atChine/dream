package v1

import (
	"github.com/gin-gonic/gin"
	"love_blog/model"
	"love_blog/utils/errmsg"
	"net/http"
)

// AddComment 新增评论
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	code := model.AddComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCommentCount 获取文章评论数量
func GetCommentCount(c *gin.Context) {
	artId := c.Param("artId")
	count := model.GetCommentCount(artId)
	c.JSON(http.StatusOK, gin.H{
		"total": count,
	})
}
