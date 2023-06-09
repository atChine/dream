package v1

import (
	"dream/model"
	"dream/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var code int
	formData, code = model.CheckLogin(&formData)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
	})
}
