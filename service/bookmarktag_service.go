package service

import (
	uuid "github.com/satori/go.uuid"
	"lilith/dao"
	"lilith/model/entity"
	"log"
	"time"
)

type BookmarkTagService struct {
	dao *dao.BookmarkTagDao
}

func NewBookmarkTagService() *BookmarkTagService {
	service := &BookmarkTagService{
		dao: dao.NewBookmarkTagDao(),
	}
	return service
}

func (b *BookmarkTagService) Create(tag entity.BookmarkTagEntity) int {
	tag.Id = uuid.NewV4().String()
	tag.CreateTime = time.Now().UnixNano() / 1e6
	i, e := b.dao.Create(tag)
	if e != nil {
		log.Println("create bookmark tag error......")
	}
	return i
}
