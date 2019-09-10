package service

import (
	uuid "github.com/satori/go.uuid"
	"lilith/dao"
	"lilith/model/dto"
	"lilith/model/entity"
	"log"
	"time"
)

type CommentService struct {
	dao *dao.CommentDao
}

func NewCommentService() *CommentService {
	service := &CommentService{
		dao: dao.NewCommentDao(),
	}
	return service
}

func (c *CommentService) CreateCommnet(entity entity.CommentEntity) error {
	entity.Id = uuid.NewV4().String()
	entity.CreateTime = time.Now().UnixNano() / 1e6
	_, e := c.dao.CreateCommnet(entity)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func (c *CommentService) CommentList(resourceId string, start, size int) ([]dto.CommentListDto, error) {
	listDto, e := c.dao.CommentList(resourceId, start, size)
	if e != nil {
		log.Println(e)
	}
	return listDto, e
}

func (c *CommentService) CountCommentByResourceId(resourceId string) int {
	return c.dao.CountCommentByResourceId(resourceId)
}

func (c *CommentService) IsCommentByUserId(resourceId, userId string) bool {
	i := c.dao.IsCommentByUserId(resourceId, userId)
	return i > 0
}
