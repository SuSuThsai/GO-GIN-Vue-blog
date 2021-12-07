package v1

import (
	"GO-GIN-Vue-blog/middleware"
	"GO-GIN-Vue-blog/model"
	"GO-GIN-Vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	c.ShouldBind(&data)
	var code int
	var token string

	_, code = model.ValidateLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		token, _ = middleware.CreatToken(data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
