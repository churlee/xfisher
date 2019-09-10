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

func (s *Spider) V2exSpider(wg *sync.WaitGroup) {
	url := "https://www.v2ex.com/?tab=hot"
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
		log.Println("v2ex response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.V2EX)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := make([]entity.FishEntity, 0)
	doc.Find("#Main").Find(".item").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i + 1
		text := selection.Find(".item_title a").Text()
		url, _ := selection.Find(".item_title a").Attr("href")
		view := selection.Find(".count_livid").Text()
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = "https://www.v2ex.com" + url
		fishEntity.View = view
		fishEntity.Type = common.V2EX
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		list = append(list, *fishEntity)
		//s.service.Create(fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.V2EX, string(bytes), 0).Err()
}
