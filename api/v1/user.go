package v1

import (
	"github.com/gin-gonic/gin"
	"love_blog/model"
	"love_blog/utils/errmsg"
	"love_blog/utils/validator"
	"net/http"
	"strconv"
)

// AddUser 添加新用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)
	msg, validCode = validator.Validate(&data)
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
	// 检查重复
	code := model.CheckUser(data.UserName)
	if code == errmsg.SUCCSE {
		model.AddUser(&data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetUserInfoById 查询单个用户的详细信息
func GetUserInfoById(c *gin.Context) {
	userId := c.Param("userId")
	userInfo, code := model.GetUserInfoById(userId)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    userInfo,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑个人信息
func EditUser(c *gin.Context) {
	userId := c.Param("userId")
	userName := c.Param("userName")
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 更新检查
	code := model.CheckEditUser(userId, userName)
	if code == errmsg.SUCCSE {
		model.EditUser(userId, &data)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetUserList 查询用户列表 + 模糊查询
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	userName := c.Query("userName")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUserList(userName, pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	userId := c.Param("userId")
	code := model.DelUser(userId)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// ChangePassword 修改密码
func ChangePassword(c *gin.Context) {
	var user model.User
	userId := c.Param("userId")
	_ = c.ShouldBindJSON(&user)
	code := model.ChangePassword(userId, &user)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
