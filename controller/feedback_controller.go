package controller

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"lilith/common"
	"lilith/model/dto"
	"lilith/service"
	"net/http"
)

var (
	feedbackService *service.FeedbackService
)

func init() {
	feedbackService = service.NewFeedbackService()
}

func CreateFeedback(c *gin.Context) {
	response := dto.Response{C: c}
	var createDto dto.CreateFeedbackDto
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
	feedbackEntity, e := feedbackService.Create(createDto, userId.(string))
	if e != nil {
		response.Error(http.StatusInternalServerError, common.CreateFeedbackError)
		return
	}
	response.Success(http.StatusOK, feedbackEntity.Id)
}
