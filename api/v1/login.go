package v1

import (
	"GO-GIN-Vue-blog/middleware"
	"GO-GIN-Vue-blog/model"
	"GO-GIN-Vue-blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login 后台登录
func Login(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)
	var code int
	var token string

	_, code = model.ValidateLogin(data.Username, data.Password)
	if code == errmsg.SUCCESS {
		setToken(c, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data.Username,
			"id":      data.ID,
			"token":   token,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// LoginFront 前台登录
func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBind(&formData)
	var code int
	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"role":    formData.Role,
		"message": errmsg.GetErrMsg(code),
	})
}

//token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "Yamada",
		},
	}
	tokenString, err := j.CreatToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   tokenString,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    user.Username,
		"id":      user.ID,
		"token":   tokenString,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
	return
}
