package v1

import (
	"GO-GIN-Vue-blog/model"
	"GO-GIN-Vue-blog/utils/errmsg"
	"GO-GIN-Vue-blog/utils/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//// UserExist 查询用户是否存在
//func UserExist(c *gin.Context) {
//
//}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var code int
	_ = c.ShouldBind(&data)
	msg, code = validate.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  code,
			"message": msg,
		})
		c.Abort()
		return
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreatUser(&data)
	} else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)

	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	//else if code == errmsg.ERROR_USERNAME_USED {
	//	c.Abort()
	//}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
