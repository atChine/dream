package v1

import (
	"dream/model"
	"dream/utils"
	"dream/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserById 查询单个用户
func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUserById(id)
	maps["userName"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetUsers 搜索用户
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	userName := c.Query("username")
	pageSize, pageNum = utils.HandlePageSizeAndPageNum(pageSize, pageNum)
	data, code, total := model.GetUsers(pageSize, pageNum, userName)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
