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
	commentService *service.CommentService
)

func init() {
	commentService = service.NewCommentService()
}

func CreateComment(c *gin.Context) {
	response := dto.Response{C: c}
	var commentDto dto.CreateCommentDto
	_ = c.BindJSON(&commentDto)
	validate := validator.New()
	e := validate.Struct(commentDto)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	userId, _ := c.Get("userId")
	commentEntity := entity.CommentEntity{
		ResourceId: commentDto.ResourceId,
		UserId:     userId.(string),
		Content:    commentDto.Content,
	}
	e = commentService.CreateCommnet(commentEntity)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.CreateCommentError)
		return
	}
	response.Success(http.StatusOK, commentEntity.Id)
}

func CommentList(c *gin.Context) {
	response := dto.Response{C: c}
	resourceId := c.Param("id")
	page := c.Param("page")
	if resourceId == "" {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	p, e := strconv.Atoi(page)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}

	total := commentService.CountCommentByResourceId(resourceId)
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
	listDto, e := commentService.CommentList(resourceId, start, size)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.ListCommentError)
		return
	}
	totalPage := int((total + size - 1) / size)
	pageDto := dto.PageDto{
		List:  listDto,
		Total: total,
		Pages: totalPage,
	}
	response.Success(http.StatusOK, pageDto)
}
