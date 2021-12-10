package middleware

import (
	"GO-GIN-Vue-blog/utils"
	"GO-GIN-Vue-blog/utils/errmsg"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(utils.JwtKey),
	}
}

var JwtKey = []byte(utils.JwtKey)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

//定义错误
var (
	TokenExpired     error = errors.New("token已过期,请重新登录")
	TokenNotValueYet error = errors.New("此token无效,请重新登录")
	TokenMalFormed   error = errors.New("token不正确,请重新登录")
	TokenInvalid     error = errors.New("token格式错误,请重新登录")
)

// CreatToken 生成TOKEN(J*JWT)
func (j *JWT) CreatToken(claims Claims) (string, error) {
	//var J *JWT
	//expireTime := time.Now().Add(7 * 24 * time.Hour)
	//claims := &Claims{
	//	UserID: user.ID,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expireTime.Unix(),
	//		IssuedAt:  time.Now().Unix(),
	//		Issuer:    "Yamada",
	//		Subject:   "user token",
	//	},
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParserToken 验证TOKEN
func (j *JWT) ParserToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return JwtKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, nil, TokenMalFormed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, nil, TokenNotValueYet
			} else {
				return nil, nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claimsData, ok := token.Claims.(*Claims); ok && token.Valid {
			return nil, claimsData, nil
		}
		return nil, nil, TokenInvalid
	}

	return nil, nil, TokenInvalid
}

// JwtToken 中间键
func JwtToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		//获取Authorization
		tokenString := context.GetHeader("Authorization")

		//判断token 的情况
		if tokenString == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		checkToken := strings.Split(tokenString, " ")
		if len(checkToken) == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		j := NewJWT()
		//解析Token
		_, claims, err := j.ParserToken(checkToken[1])
		if err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusUnauthorized, gin.H{
					"status":  errmsg.ERROR,
					"message": TokenExpired,
					"data":    nil,
				})
				context.Abort()
				return
			}
			context.JSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
				"data":    nil,
			})
			context.Abort()
			return
		}
		context.Set("user_id", claims)
		context.Next()
	}
}
