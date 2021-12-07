package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
			AllowedHeaders:   []string{"Origin"},
			ExposedHeaders:   []string{"Content-Length", "Authorization"},
			AllowCredentials: true,
			MaxAge:           43200,
		})
	}
}
