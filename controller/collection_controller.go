package controller

import (
	"github.com/gin-gonic/gin"
	"lilith/common"
	"lilith/model/dto"
	"lilith/model/entity"
	"lilith/service"
	"net/http"
)

var (
	collectionService *service.CollectionService
)

func init() {
	collectionService = service.NewCollectionService()
}

func AddCollection(c *gin.Context) {
	response := dto.Response{C: c}
	userId, _ := c.Get("userId")
	navId := c.Param("id")
	if navId == "" {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	exist := collectionService.Exist(userId.(string), navId)
	if exist > 0 {
		response.Success(http.StatusOK, navId)
		return
	}
	collectionEntity := entity.CollectionEntity{
		UserId: userId.(string),
		NavId:  navId,
	}
	_, e := collectionService.Create(collectionEntity)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	go navService.AddCollection(navId)
	response.Success(http.StatusOK, navId)
}

func CollectionAll(c *gin.Context) {
	response := dto.Response{C: c}
	userId, _ := c.Get("userId")
	all, e := collectionService.All(userId.(string))
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ListCollectionError)
		return
	}
	response.Success(http.StatusOK, all)
}
