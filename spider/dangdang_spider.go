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

func (s *Spider) DangDangSpider(wg *sync.WaitGroup) {
	url := "http://bang.dangdang.com/books/bestsellers/01.00.00.00.00.00-recent7-0-0-1-1"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	defer wg.Done()
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("user-agent", common.UserAgent)
	req.Header.Add("Host", "bang.dangdang.com")
	req.Header.Add("Referer", "http://bang.dangdang.com/books/bestsellers/01.00.00.00.00.00-24hours-0-0-1-1")
	response, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if response.StatusCode != 200 {
		log.Println("dangdang response err...")
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
	doc.Find(".bang_list").Find("li").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i + 1
		text := selection.Find(".name a").Text() + "---"
		selection.Find(".publisher_info").Each(func(j int, js *goquery.Selection) {
			if j == 0 {
				text += js.Text()
			}
		})
		text = common.ConvertToString(text, "gbk", "utf-8")
		url, _ := selection.Find(".name a").Attr("href")
		url = common.ConvertToString(url, "gbk", "utf-8")
		view := ""
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = url
		fishEntity.View = view
		fishEntity.Type = common.DangDang
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		//s.service.Create(fishEntity)
		list = append(list, *fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.DangDang, string(bytes), 0).Err()
}
