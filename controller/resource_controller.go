package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"lilith/common"
	"lilith/model/dto"
	"lilith/model/entity"
	"lilith/service"
	"net/http"
	"strconv"
)

var (
	resourceService *service.ResourceService
)

func init() {
	resourceService = service.NewResourceService()
}

func CreateResource(c *gin.Context) {
	userId, _ := c.Get("userId")
	response := dto.Response{C: c}
	var resourceDto dto.CreateResourceDto
	e := c.BindJSON(&resourceDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	validate := validator.New()
	e = validate.Struct(resourceDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	resourceEntity := entity.ResourceEntity{
		Title:   resourceDto.Title,
		Content: resourceDto.Content,
		Type:    resourceDto.Type,
		LinkUrl: resourceDto.LinkUrl,
		LinkPwd: resourceDto.LinkPwd,
		UserId:  userId.(string),
	}
	e = resourceService.Create(resourceEntity)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.CreateResourceError)
		return
	}
	response.Success(http.StatusOK, resourceEntity.Id)
}

func ResourceList(c *gin.Context) {
	response := dto.Response{C: c}
	page := c.Query("page")
	resourceType := c.Query("type")
	title := c.Query("title")
	if page == "" {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	p, e := strconv.Atoi(page)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	total := resourceService.ResourceCount(resourceType, title)
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
	resourceListDto, e := resourceService.ResourceList(resourceType, title, start, size)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ListResourceError)
		return
	}
	totalPage := int((total + size - 1) / size)
	pageDto := dto.PageDto{
		List:  resourceListDto,
		Total: total,
		Pages: totalPage,
	}
	response.Success(http.StatusOK, pageDto)
}

func ResourceInfo(c *gin.Context) {
	response := dto.Response{C: c}
	id := c.Param("id")
	if id == "" {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userId, _ := c.Get("userId")
	go resourceService.AddView(id)
	info := resourceService.ResourceInfo(id)
	isCommentByUserId := commentService.IsCommentByUserId(id, userId.(string))
	if !isCommentByUserId {
		info.LinkUrl = ""
		info.LinkPwd = ""
	}
	info.IsComment = isCommentByUserId
	response.Success(http.StatusOK, info)
}
