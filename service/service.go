package service

import (
	"lilith/dao"
)

type Service struct {
	dao *dao.Dao
}

func NewService() *Service {
	service := &Service{
		dao: dao.NewDB(),
	}
	return service
}
