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

func (s *Spider) WeiboSpider(wg *sync.WaitGroup) {
	url := "https://s.weibo.com/top/summary?cate=realtimehot"
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
		log.Println("weibo response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.Weibo)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := make([]entity.FishEntity, 0)
	doc.Find("#pl_top_realtimehot").Find("tr").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i
		text := selection.Find(".td-02 a").Text()
		if text == "" {
			return
		}
		url, _ := selection.Find(".td-02 a").Attr("href")
		view := selection.Find(".td-02 span").Text()
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = "https://s.weibo.com" + url
		fishEntity.View = view
		fishEntity.Type = common.Weibo
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		list = append(list, *fishEntity)
		//s.service.Create(fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.Weibo, string(bytes), 0).Err()
}
