package job

import (
	"lilith/spider"
	"log"
	"net/http"
	"time"
)

func InitSpider() {
	log.Println("定时任务启动......")
	s := spider.NewSpider(&http.Client{})
	for range time.Tick(time.Minute * 10) {
		log.Println("数据同步开始......")
		s.Start()
	}
}
