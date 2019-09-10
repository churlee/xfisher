package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	CommonCache "lilith/cache"
	"lilith/common"
	"lilith/model/dto"
	"net/http"
	"time"
)

func OnlineMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := dto.Response{C: c}
		ip := c.ClientIP()
		if ip == "" {
			response.Error(http.StatusForbidden, common.ServerError)
			c.Abort()
			return
		}
		onlineCache := CommonCache.GetOnlineCache()
		_, ex := onlineCache.Get(ip)
		t := time.Now().UnixNano() / 1e6
		if ex {
			onlineCache.Delete(ip)
		}
		onlineCache.Set(ip, t, cache.DefaultExpiration)
		c.Next()
	}
}
