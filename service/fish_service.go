package service

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"lilith/common"
	"lilith/config"
	"lilith/dao"
	"lilith/model/entity"
)

//noinspection GoNameStartsWithPackageName
type FishService struct {
	dao         *dao.FishDao
	redisClient *redis.Client
}

func NewFishService() *FishService {
	service := &FishService{
		dao:         dao.NewFishDao(),
		redisClient: config.GetRedisClient(),
	}
	return service
}

func (f *FishService) Create(fishEntity *entity.FishEntity) {
	f.dao.Create(fishEntity)
}

func (f *FishService) FishAll(fishType string) ([]entity.FishEntity, error) {
	result, err := f.redisClient.Get(common.SpiderTypePrefix + fishType).Result()
	fishList := make([]entity.FishEntity, 0)
	if err != nil {
		//fishList, e := f.dao.All(fishType)
		//if e != nil {
		//	log.Println("Get v2 Error......")
		//	return nil, e
		//}
		//go func() {
		//	bytes, _ := json.Marshal(fishList)
		//	navKey := common.SpiderTypePrefix + fishType
		//	err := f.redisClient.Set(navKey, string(bytes), time.Minute*10).Err()
		//	if err != nil {
		//		fmt.Println("FishAll redis set error...")
		//	}
		//}()
		return fishList, nil
	}
	_ = json.Unmarshal([]byte(result), &fishList)
	return fishList, nil

}

func (f *FishService) DeleteByType(fishType string) {
	f.dao.DeleteByType(fishType)
}
