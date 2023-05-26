package v1

import (
	"github.com/gin-gonic/gin"
	"love_blog/model"
	"love_blog/utils/errmsg"
	"net/http"
	"strconv"
)

// AddCate 新增分类标签
func AddCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBindJSON(&cate)
	// 查看分类标签是否存在
	code := model.CheckCate(cate.Name)
	if code == errmsg.SUCCSE {
		model.AddCate(&cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DelCateById 根据id删除分类标签
func DelCateById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errmsg.BadRequest(c, "输入的id不合法")
	}
	code := model.DelCateById(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditCateById 根据id编辑分类名字
func EditCateById(c *gin.Context) {
	var cate model.Category
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errmsg.BadRequest(c, "输入的id不合法")
	}
	_ = c.ShouldBindJSON(&cate)
	code := model.CheckCate(cate.Name)
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	if code == errmsg.SUCCSE {
		code = model.EditCateById(id, &cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateInfoById 通过id查询单个分类详细信息
func GetCateInfoById(c *gin.Context) {
	id, err := strconv.Atoi("id")
	if err != nil {
		errmsg.BadRequest(c, "输入的id不合法")
	}
	data, code := model.GetCateInfoById(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetCateList 查询分类列表
func GetCateList(c *gin.Context) {
	pageSize, _ := strconv.Atoi("pageSize")
	pageNum, _ := strconv.Atoi("pageNUm")
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total, code := model.GetCateList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
