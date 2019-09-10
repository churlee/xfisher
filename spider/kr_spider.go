package spider

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	uuid "github.com/satori/go.uuid"
	"lilith/common"
	"lilith/model/entity"
	"log"
	"net/http"
	"sync"
	"time"
)

func (s *Spider) krSpider(wg *sync.WaitGroup) {
	url := "https://36kr.com/information/web_news"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	defer wg.Done()
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("user-agent", common.UserAgent)
	response, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if response.StatusCode != 200 {
		log.Println("36kr response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.KR)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := make([]entity.FishEntity, 0)
	doc.Find(".information-flow-list").Find(".information-flow-item").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i + 1
		text := selection.Find(".article-item-title").Text()
		url, _ := selection.Find(".article-item-title").Attr("href")
		view := selection.Find(".article-item-channel").Text()
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = "https://36kr.com" + url
		fishEntity.View = view
		fishEntity.Type = common.KR
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		//s.service.Create(fishEntity)
		list = append(list, *fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.KR, string(bytes), 0).Err()
	//_ = s.redisClient.Del(common.SpiderTypePrefix + common.KR).Err()
}
