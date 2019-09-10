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

func (s *Spider) WeixinSpider(wg *sync.WaitGroup) {
	url := "http://www.gsdata.cn/rank/wxarc"
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
		log.Println("weixin response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.WeiXin)
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	list := make([]entity.FishEntity, 0)
	doc.Find("#rank_data").Find("tr").Each(func(i int, selection *goquery.Selection) {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = i
		text := selection.Find(".tdhidden a").Text()
		if text == "" {
			return
		}
		url, _ := selection.Find(".tdhidden a").Attr("href")
		view := ""
		selection.Find("td").Each(func(j int, js *goquery.Selection) {
			if j == 1 {
				text += "   《" + js.Text() + "》"
			}
			if j == 3 {
				view = js.Text()
			}
		})
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = url
		fishEntity.View = view
		fishEntity.Type = common.WeiXin
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		list = append(list, *fishEntity)
		//s.service.Create(fishEntity)
	})
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.WeiXin, string(bytes), 0).Err()
}
