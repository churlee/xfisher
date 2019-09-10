package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"lilith/dao"
	"lilith/model/dto"
	"lilith/model/entity"
	"log"
	"time"
)

type FeedbackService struct {
	dao *dao.FeedbackDao
}

func NewFeedbackService() *FeedbackService {
	feedbackService := &FeedbackService{
		dao: dao.NewFeedbackDao(),
	}
	return feedbackService
}

func (f *FeedbackService) Create(fd dto.CreateFeedbackDto, userId string) (entity.FeedbackEntity, error) {
	feedbackEntity := entity.FeedbackEntity{
		Id:         uuid.NewV4().String(),
		Content:    fd.Content,
		UserId:     userId,
		CreateTime: time.Now().UnixNano() / 1e6,
	}
	i := f.dao.Create(feedbackEntity)
	if i != 1 {
		log.Println("create feedback error...")
		return feedbackEntity, fmt.Errorf("create feedback error...")
	}
	return feedbackEntity, nil
}
