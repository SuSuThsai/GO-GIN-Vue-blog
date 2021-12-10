package v1

import (
	"GO-GIN-Vue-blog/model"
	"GO-GIN-Vue-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetProfile 获取个人信息设置
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	profile, code := model.GetProfile(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    profile,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateProfile 更新个人设置
func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.UpdateProfile(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
