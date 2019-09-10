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

type CommunicationService struct {
	dao *dao.CommunicationDao
}

func NewCommunicationService() *CommunicationService {
	service := &CommunicationService{
		dao: dao.NewCommunicationDao(),
	}
	return service
}

func (s *CommunicationService) Create(d dto.CreateCommunicationDto, userId string) (entity.CommunicationEntity, error) {
	e := entity.CommunicationEntity{
		Id:         uuid.NewV4().String(),
		Title:      d.Title,
		Node:       d.Node,
		Content:    d.Content,
		View:       0,
		UserId:     userId,
		CreateTime: time.Now().UnixNano() / 1e6,
	}
	i := s.dao.Create(e)
	if i != 1 {
		log.Println("create communication error...")
		return e, fmt.Errorf("")
	}
	return e, nil
}

func (s *CommunicationService) ListByNode(node, title string, start, size int) ([]dto.ListCommunicationDto, error) {
	listCommunicationDto, e := s.dao.ListByNode(node, title, start, size)
	if e != nil {
		log.Println("list by node error...")
	}
	return listCommunicationDto, e
}

func (s *CommunicationService) CountByNode(node, title string) int {
	return s.dao.CountByNode(node, title)
}

func (s *CommunicationService) CommunicationInfo(id string) dto.CommunicationInfoDto {
	return s.dao.FindById(id)
}

func (s *CommunicationService) AddView(id string) {
	i := s.dao.AddView(id)
	if i != 1 {
		log.Println("add communication view error...")
	}
}
