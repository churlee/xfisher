package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"lilith/common"
	"lilith/model/dto"
	"lilith/model/entity"
	"lilith/service"
	"net/http"
)

var (
	navService *service.NavService
)

func init() {
	navService = service.NewNavService()
}

func NavList(c *gin.Context) {
	response := dto.Response{C: c}
	navType := c.Param("type")
	if navType == "" {
		response.Success(http.StatusOK, make([]interface{}, 0))
	}
	navAll, err := navService.NavListByType(navType)
	if err != nil {
		response.Error(http.StatusInternalServerError, common.GetNavAllError)
	}
	response.Success(http.StatusOK, navAll)
}

func NavHot(c *gin.Context) {
	response := dto.Response{C: c}
	navAll, err := navService.NavHot()
	if err != nil {
		response.Error(http.StatusInternalServerError, common.GetNavAllError)
	}
	response.Success(http.StatusOK, navAll)
}

func NavLike(c *gin.Context) {
	response := dto.Response{C: c}
	id := c.Param("id")
	navService.AddNavLike(id)
	response.Success(http.StatusOK, id)
}

func NavGo(c *gin.Context) {
	url := c.Query("url")
	id := c.Query("id")
	if url == "" || id == "" {
		c.Redirect(http.StatusMovedPermanently, "/error")
		return
	}
	go navService.AddNavView(id)
	c.Redirect(http.StatusMovedPermanently, url)
}

func CreateNav(c *gin.Context) {
	response := dto.Response{C: c}
	var navEntity entity.NavEntity
	err := c.BindJSON(&navEntity)
	if err != nil {
		response.Error(http.StatusBadRequest, common.CreateNavError)
		return
	}
	uid := uuid.NewV4()
	navEntity.Id = uid.String()
	navService.Create(navEntity)
}

func FindByTitle(c *gin.Context) {
	response := dto.Response{C: c}
	navTitle := c.Param("title")
	if navTitle == "" {
		response.Success(http.StatusOK, make([]interface{}, 0))
	}
	navList, err := navService.FindByTitle(navTitle)
	if err != nil {
		response.Error(http.StatusInternalServerError, common.GetNavAllError)
	}
	response.Success(http.StatusOK, navList)
}
