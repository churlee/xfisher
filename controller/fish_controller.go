package controller

import (
	"github.com/gin-gonic/gin"
	"lilith/common"
	"lilith/model/dto"
	"lilith/service"
	"net/http"
)

var (
	fishService *service.FishService
)

func init() {
	fishService = service.NewFishService()
}

//noinspection GoNilness
func FishAll(c *gin.Context) {
	response := dto.Response{C: c}
	fishType := c.Param("type")
	fishAll, err := fishService.FishAll(fishType)
	if err != nil {
		response.Error(http.StatusInternalServerError, common.GetNavAllError)
	}
	response.Success(http.StatusOK, fishAll)
}
