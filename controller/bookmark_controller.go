package controller

import (
	"encoding/json"
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
	bookmarkService *service.BookmarkService
)

func init() {
	bookmarkService = service.NewBookmarkService()
}

func CreateBookmark(c *gin.Context) {
	response := dto.Response{C: c}
	var createDto dto.CreateBookmarkDto
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
	share, _ := strconv.Atoi(createDto.Share)
	tBytes, err := json.Marshal(createDto.Tags)
	if err != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	tags := string(tBytes)
	if len(tags) > 250 {
		response.Error(http.StatusBadRequest, common.TagsTooLongError)
		return
	}
	createEntity := entity.BookmarkEntity{
		Title:  createDto.Title,
		Url:    createDto.Url,
		Share:  share,
		Tags:   tags,
		UserId: userId.(string),
	}
	bookmarkEntity, e := bookmarkService.Create(createEntity)
	if e != nil {
		response.Error(http.StatusInternalServerError, common.CreateCommunicationError)
		return
	}
	response.Success(http.StatusOK, bookmarkEntity.Id)
}

func ListBookmark(c *gin.Context) {
	response := dto.Response{C: c}
	page := c.Query("page")
	bookmarkType := c.Query("type")
	title := c.Query("title")
	p, e := strconv.Atoi(page)
	if e != nil {
		response.Error(http.StatusBadRequest, common.ParamsError)
		return
	}
	me := bookmarkType == "me"
	uid := ""
	if me {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.Error(http.StatusUnauthorized, common.AuthError)
			return
		}
		userId, ok := common.CheckToken(token)
		if !ok {
			response.Error(http.StatusUnauthorized, common.AuthError)
			return
		}
		uid = userId
	}
	total := bookmarkService.Count(title, uid)
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
	listCommunicationDto, e := bookmarkService.List(title, uid, start, size)
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
