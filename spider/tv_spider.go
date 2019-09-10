package spider

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"lilith/common"
	"lilith/model/entity"
	"log"
	"net/http"
	"sync"
	"time"
)

func (s *Spider) TvSpider(wg *sync.WaitGroup) {
	url := "https://movie.douban.com/j/search_subjects?type=tv&tag=热门&sort=recommend&page_limit=20&page_start=0"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	defer wg.Done()
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("user-agent", common.UserAgent)
	req.Header.Add("Host", "movie.douban.com")
	response, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if response.StatusCode != 200 {
		log.Println("tv response err...")
		return
	}
	defer response.Body.Close()
	//s.service.DeleteByType(common.BiliBili)
	body, _ := ioutil.ReadAll(response.Body)
	var tvDto TvDto
	err = json.Unmarshal(body, &tvDto)
	if err != nil {
		log.Println("tv json error...")
		return
	}
	list := make([]entity.FishEntity, 0)
	for _, v := range tvDto.Subjects {
		fishEntity := &entity.FishEntity{}
		fishEntity.Order = 0
		text := v.Title
		url := v.Url
		view := v.Rate
		fishEntity.Title = text
		fishEntity.Id = uuid.NewV4().String()
		fishEntity.Url = url
		fishEntity.View = view
		fishEntity.Type = common.Tv
		fishEntity.CreateTime = time.Now().UnixNano() / 1e6
		//s.service.Create(fishEntity)
		list = append(list, *fishEntity)
	}
	bytes, _ := json.Marshal(list)
	_ = s.redisClient.Set(common.SpiderTypePrefix+common.Tv, string(bytes), 0).Err()
}

type TvDto struct {
	Subjects []TvList
}

type TvList struct {
	Cover    string `json:"cover"`
	CoverX   int    `json:"cover_x"`
	CoverY   int    `json:"cover_y"`
	Id       string `json:"id"`
	IsNew    bool   `json:"is_new"`
	Playable bool   `json:"playable"`
	Rate     string `json:"rate"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}
