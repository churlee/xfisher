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

func (s *Spider) GuoKeSpider(wg *sync.WaitGroup) {
	url := "https://www.guokr.com/scientific/"
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
		log.Println("guoke response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.BiliBili)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := make([]entity.FishEntity, 0)
	doc.Find(".article").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i + 1
		text := selection.Find("h3").Find("a").Text()
		if text == "" {
			return
		}
		url, _ := selection.Find("h3").Find("a").Attr("href")
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = url
		fishEntity.Type = common.GuoKe
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		//s.service.Create(fishEntity)
		list = append(list, *fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.GuoKe, string(bytes), 0).Err()
}
