package controller

import (
	"github.com/gin-gonic/gin"
	CommonCache "lilith/cache"
	"lilith/model/dto"
	"net/http"
)

func OnlineUsersNo(c *gin.Context) {
	response := dto.Response{C: c}
	onlineCache := CommonCache.GetOnlineCache()
	count := onlineCache.ItemCount()
	response.Success(http.StatusOK, count)
}
