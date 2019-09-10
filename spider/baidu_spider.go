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

func (s *Spider) BaiduSpider(wg *sync.WaitGroup) {
	url := "http://top.baidu.com/buzz?b=1"
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
		log.Println("baidu response err...")
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
	doc.Find(".list-table").Find("tr").Each(func(i int, selection *goquery.Selection) {
		c, exists := selection.Attr("class")
		if exists {
			if c != "hideline" {
				return
			}
		}
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i
		text := selection.Find(".list-title").Text()
		text = common.ConvertToString(text, "gbk", "utf-8")
		if text == "" {
			return
		}
		url, _ := selection.Find(".list-title").Attr("href")
		url = common.ConvertToString(url, "gbk", "utf-8")
		view := selection.Find(".last span").Text()
		//view = common.ConvertToString(view, "gbk", "utf-8")
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = url
		fishEntity.View = view
		fishEntity.Type = common.Baidu
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		list = append(list, *fishEntity)
		//s.service.Create(fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.Baidu, string(bytes), 0).Err()
}
