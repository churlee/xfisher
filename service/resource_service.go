package service

import (
	uuid "github.com/satori/go.uuid"
	"lilith/dao"
	"lilith/model/dto"
	"lilith/model/entity"
	"log"
	"time"
)

type ResourceService struct {
	dao *dao.ResourceDao
}

func NewResourceService() *ResourceService {
	service := &ResourceService{
		dao: dao.NewResourceDao(),
	}
	return service
}

func (s *ResourceService) Create(entity entity.ResourceEntity) error {
	entity.Id = uuid.NewV4().String()
	entity.CreateTime = time.Now().UnixNano() / 1e6
	entity.View = 0
	_, e := s.dao.Create(entity)
	if e != nil {
		log.Println(e)
		return e
	}
	return nil
}

func (s *ResourceService) ResourceCount(resourceType, title string) int {
	return s.dao.ResourceCount(resourceType, title)
}

func (s *ResourceService) ResourceList(resourceType, title string, start, size int) ([]dto.ResourceListDto, error) {
	return s.dao.ResourceList(resourceType, title, start, size)
}

func (s *ResourceService) ResourceInfo(id string) dto.ResourceInfoDto {
	return s.dao.ResourceInfo(id)
}

func (s *ResourceService) AddView(id string) {
	i := s.dao.AddView(id)
	if i != 1 {
		log.Println("add resource view error...")
	}
}
