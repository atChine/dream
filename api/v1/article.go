package v1

import (
	"github.com/gin-gonic/gin"
	"love_blog/model"
	"love_blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddArt AddArticle 新增文章
func AddArt(c *gin.Context) {
	var articleData model.Article
	_ = c.ShouldBindJSON(&articleData)
	code := model.AddArt(&articleData)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articleData,
		"message": errmsg.GetErrMsg(code),
	})
}

// DelArtById 通过id删除指定文章
func DelArtById(c *gin.Context) {
	id := c.Param("artId")
	artId, err := strconv.Atoi(id)
	if err != nil {
		errmsg.BadRequest(c, "输入的id不合法")
	}
	code := model.DelArtById(artId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EdiArtById 根据id修改文章
func EdiArtById(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	artId, err := strconv.Atoi(c.Param("artId"))
	if err != nil {
		errmsg.BadRequest(c, "输入id不合法")
	}
	code := model.EdiArtById(artId, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfoById 根据id查询文章详情
func GetArtInfoById(c *gin.Context) {
	artId, _ := strconv.Atoi(c.Param("artId"))
	data, code := model.GetArtInfoById(artId)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfoByCate 分页查询分类下的所有文章
func GetArtInfoByCate(c *gin.Context) {
	cateId, _ := strconv.Atoi(c.Param("cateId"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize < 0:
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	cateArtList, code, total := model.GetArtInfoByCate(cateId, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cateArtList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 分页查询所有文章/文章标题搜索
func GetArtInfo(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	artTitle := c.Query("artTitle")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	// 全部查询
	if len(artTitle) == 0 {
		artList, code, total := model.GetArtInfo(pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    artList,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	// 按照标题搜索文章
	artList, code, total := model.GetArtInfoByTitle(pageSize, pageNum, artTitle)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    artList,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
