package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"lilith/common"
	"lilith/config"
	"lilith/dao"
	"lilith/model/entity"
	"log"
	"time"
)

//noinspection GoNameStartsWithPackageName
type NavService struct {
	dao         *dao.NavDao
	redisClient *redis.Client
}

func NewNavService() *NavService {
	service := &NavService{
		dao:         dao.NewNavDao(),
		redisClient: config.GetRedisClient(),
	}
	return service
}

func (s *NavService) Create(nav entity.NavEntity) {
	i := s.dao.Create(nav)
	if i == 1 {
		log.Println("nav insert success...")
	}
}

func (s *NavService) NavListByType(navType string) ([]entity.NavEntity, error) {
	result, err := s.redisClient.Get(common.NavTypePrefix + navType).Result()
	if err != nil {
		navEntities, err := s.dao.NavListByType(navType)
		if err != nil {
			log.Println("Get NavAll Error......")
			return nil, err
		}
		go func() {
			bytes, _ := json.Marshal(navEntities)
			navKey := common.NavTypePrefix + navType
			err := s.redisClient.Set(navKey, string(bytes), time.Minute*10).Err()
			if err != nil {
				fmt.Println("NavAll redis set error...")
			}
		}()
		return navEntities, nil
	}
	navAll := make([]entity.NavEntity, 0)
	_ = json.Unmarshal([]byte(result), &navAll)
	return navAll, nil
}

func (s *NavService) NavHot() ([]entity.NavEntity, error) {
	result, err := s.redisClient.Get(common.NavTypePrefix + "hot").Result()
	if err != nil {
		navEntities, err := s.dao.NavHot()
		if err != nil {
			log.Println("Get NavAll Error......")
			return nil, err
		}
		go func() {
			bytes, _ := json.Marshal(navEntities)
			navKey := common.NavTypePrefix + "hot"
			err := s.redisClient.Set(navKey, string(bytes), time.Minute*10).Err()
			if err != nil {
				fmt.Println("NavAll redis set error...")
			}
		}()
		return navEntities, nil
	}
	navAll := make([]entity.NavEntity, 0)
	_ = json.Unmarshal([]byte(result), &navAll)
	return navAll, nil
}

func (s *NavService) AddNavView(id string) {
	s.dao.AddNavView(id)
}

func (s *NavService) AddCollection(id string) {
	s.dao.AddCollection(id)
}

func (s *NavService) AddNavLike(id string) {
	s.dao.AddLike(id)
}

func (s *NavService) DeleteNavFromRedis(key string) {
	i, _ := s.redisClient.Del(key).Result()
	if i == 1 {
		log.Println("delete redis key---" + key + " success...")
	}
}

func (s *NavService) FindByTitle(title string) ([]entity.NavEntity, error) {
	navEntities, err := s.dao.FindByTitle(title)
	if err != nil {
		log.Println("Find Nav By Title Error......")
		return nil, err
	}
	return navEntities, nil
}
