package v1

import (
	"dream/model"
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
