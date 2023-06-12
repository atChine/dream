package v1

import (
	"dream/model"
	"dream/utils"
	"dream/utils/errmsg"
	"dream/utils/validator"
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

// AddUser 增加用户
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	msg, validCode := validator.Validate(user)
	if validCode != errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}
	// 查询username在不在
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCSE {
		model.AddUser(&user)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)
	code := model.CheckUpUser(id, user.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &user)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
