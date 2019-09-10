package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"lilith/common"
	"lilith/model/dto"
	"lilith/service"
	"net/http"
	"strconv"
)

var (
	communicationService *service.CommunicationService
)

func init() {
	communicationService = service.NewCommunicationService()
}

func CreateCommunication(c *gin.Context) {
	response := dto.Response{C: c}
	var createDto dto.CreateCommunicationDto
	e := c.Bind(&createDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(createDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userId, _ := c.Get("userId")
	communicationEntity, e := communicationService.Create(createDto, userId.(string))
	if e != nil {
		response.Error(http.StatusInternalServerError, common.CreateCommunicationError)
		return
	}
	response.Success(http.StatusOK, communicationEntity.Id)
}

func ListByNode(c *gin.Context) {
	response := dto.Response{C: c}
	node := c.Query("node")
	page := c.Query("page")
	title := c.Query("title")
	p, e := strconv.Atoi(page)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	total := communicationService.CountByNode(node, title)
	if total == 0 {
		pageDto := dto.PageDto{
			List:  make([]interface{}, 0),
			Total: 0,
			Pages: 0,
		}
		response.Success(http.StatusOK, pageDto)
		return
	}
	if p < 1 {
		p = 1
	}
	size := 20
	start := (p - 1) * size
	listCommunicationDto, e := communicationService.ListByNode(node, title, start, size)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ListCommunicationError)
		return
	}
	totalPage := int((total + size - 1) / size)
	pageDto := dto.PageDto{
		List:  listCommunicationDto,
		Total: total,
		Pages: totalPage,
	}
	response.Success(http.StatusOK, pageDto)
}

func CommunicationInfo(c *gin.Context) {
	response := dto.Response{C: c}
	id := c.Param("id")
	if id == "" {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	go communicationService.AddView(id)
	info := communicationService.CommunicationInfo(id)
	response.Success(http.StatusOK, info)
}
