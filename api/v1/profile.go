package v1

import (
	"dream/model"
	"dream/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProfileById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfileById(id)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
