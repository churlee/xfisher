package middleware

import (
	"github.com/gin-gonic/gin"
	"lilith/common"
	"lilith/model/dto"
	"net/http"
)

func AuthJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := dto.Response{C: c}
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Error(http.StatusUnauthorized, common.AuthError)
			c.Abort()
			return
		}
		userId, ok := common.CheckToken(token)
		if !ok {
			response.Error(http.StatusUnauthorized, common.AuthError)
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
