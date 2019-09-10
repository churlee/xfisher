package service

import (
	"lilith/dao"
	"lilith/model/entity"
	"log"
)

type CollectionService struct {
	dao *dao.CollectionDao
}

func NewCollectionService() *CollectionService {
	service := &CollectionService{
		dao: dao.NewCollectionDao(),
	}
	return service
}

func (c *CollectionService) Create(cEntity entity.CollectionEntity) (int, error) {
	return c.dao.Create(cEntity)
}

func (c *CollectionService) Exist(userId, navId string) int {
	return c.dao.Exist(userId, navId)
}

func (c *CollectionService) All(userId string) ([]entity.NavEntity, error) {
	all, e := c.dao.All(userId)
	if e != nil {
		log.Println(e)
	}
	return all, e
}
