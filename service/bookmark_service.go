package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"lilith/dao"
	"lilith/model/dto"
	"lilith/model/entity"
	"log"
	"time"
)

type BookmarkService struct {
	dao *dao.BookmarkDao
}

func NewBookmarkService() *BookmarkService {
	service := &BookmarkService{
		dao: dao.NewBookmarkDao(),
	}
	return service
}

func (b *BookmarkService) Create(bookmark entity.BookmarkEntity) (entity.BookmarkEntity, error) {
	bookmark.Id = uuid.NewV4().String()
	bookmark.CreateTime = time.Now().UnixNano() / 1e6
	i, e := b.dao.Create(bookmark)
	if e != nil {
		log.Println("create bookmark error......")
	}
	if i != 1 {
		e = errors.New("create bookmark error")
	}
	return bookmark, e
}

func (b *BookmarkService) List(title, uid string, start, size int) ([]dto.BookmarkListDto, error) {
	return b.dao.List(title, uid, start, size)
}

func (b *BookmarkService) Count(title, uid string) int {
	return b.dao.Count(title, uid)
}
