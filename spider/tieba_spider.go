package spider

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"lilith/common"
	"log"
	"net/http"
	"sync"
)

func (s *Spider) TiebaSpider(wg *sync.WaitGroup) {
	url := "http://tieba.baidu.com/hottopic/browse/topicList"
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
		log.Println("tieba response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.BiliBili)
	body, _ := ioutil.ReadAll(response.Body)
	js, err2 := simplejson.NewJson(body)
	if err2 != nil {
		log.Println("tieba response err...")
		return
	}
	arrayStr, err := js.Get("data").Get("bang_topic").Get("topic_list").String()
	//doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(err)
		return
	}
	tiebaList := make([]TieBaDto, 0)
	err = json.Unmarshal([]byte(arrayStr), &tiebaList)
	if err != nil {
		log.Println(err)
		return
	}
	//list := make([]entity.FishEntity, 0)
	log.Println(len(tiebaList))
	//for t := range tiebaList {
	//	fishEntity := &entity.FishEntity{}
	//	text := t.
	//	url, _ := selection.Find(".title").Attr("href")
	//	view := selection.Find(".pts div").Text()
	//	fishEntity.Title = text
	//	fishEntity.Id = uuid.NewV4().String()
	//	fishEntity.Url = url
	//	fishEntity.View = view
	//	fishEntity.Type = common.BiliBili
	//	fishEntity.CreateTime = time.Now().UnixNano() / 1e6
	//	//s.service.Create(fishEntity)
	//	list = append(list, *fishEntity)
	//}
	//bytes, _ := json.Marshal(list)
	//_ = s.redisClient.Set(common.SpiderTypePrefix+common.BiliBili, string(bytes), 0).Err()
}

type TieBaDto struct {
	TopicId            int    `json:"topic_id"`
	TopicName          string `json:"topic_name"`
	TopicDesc          string `json:"topic_desc"`
	Abstract           string `json:"abstract"`
	TopicPic           string `json:"topic_pic"`
	Tag                int    `json:"tag"`
	DiscussNum         int    `json:"discuss_num"`
	IdxNum             int    `json:"idx_num"`
	CreateTime         int    `json:"create_time"`
	ContentNum         int    `json:"content_num"`
	TopicAvatar        string `json:"topic_avatar"`
	TopicUrl           string `json:"topic_url"`
	TopicDefaultAvatar string `json:"topic_default_avatar"`
}
