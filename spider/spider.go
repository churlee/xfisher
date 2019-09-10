package spider

import (
	"github.com/go-redis/redis"
	"lilith/config"
	"lilith/service"
	"log"
	"net/http"
	"sync"
)

type Spider struct {
	client      *http.Client
	service     *service.FishService
	redisClient *redis.Client
}

func NewSpider(client *http.Client) *Spider {
	spider := &Spider{
		client:      client,
		service:     service.NewFishService(),
		redisClient: config.GetRedisClient(),
	}
	return spider
}

func (s *Spider) Start() {
	var wg sync.WaitGroup
	wg.Add(14)
	go s.V2exSpider(&wg)
	go s.krSpider(&wg)
	go s.BilibiliSpider(&wg)
	go s.WeiboSpider(&wg)
	go s.BaiduSpider(&wg)
	go s.WeixinSpider(&wg)
	go s.ZhihuSpider(&wg)
	go s.QiDianSpider(&wg)
	go s.DangDangSpider(&wg)
	go s.TvSpider(&wg)
	go s.ItHomeSpider(&wg)
	go s.MovieSpider(&wg)
	go s.GithubSpider(&wg)
	go s.GuoKeSpider(&wg)
	wg.Wait()
	log.Println("数据同步完成......")
	//_ = s.redisClient.Del(common.SpiderTypePrefix + common.FishAll).Err()
}
