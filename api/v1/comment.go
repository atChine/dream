package v1

import (
	"github.com/gin-gonic/gin"
	"love_blog/model"
	"love_blog/utils/errmsg"
	"net/http"
	"strconv"
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

// GetCommentById 获取单个评论信息
func GetCommentById(c *gin.Context) {
	comId := c.Param("com_id")
	comInfo, code := model.GetCommentById(comId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data:":   comInfo,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtCommentList 展示页面获取评论列表
func GetArtCommentList(c *gin.Context) {
	artId := c.Param("artId")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	commentList, total, code := model.GetArtCommentList(artId, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    commentList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteCommentById 删除评论
func DeleteCommentById(c *gin.Context) {
	comId := c.Param("comId")
	code := model.DeleteCommentById(comId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCommentList 后台获取所有评论列表
func GetCommentList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := model.GetCommentList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})

}
