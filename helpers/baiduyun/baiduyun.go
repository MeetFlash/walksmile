package baiduyun

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"walksmile/controllers"
	"walksmile/models"

	"encoding/json"
)

func Fetcher() {

	auto := newFetcher()
	for _, v := range auto {
		category := controllers.Catagory_caculate(v.Title)
		link := &models.Link{}
		link.FkUK = int(v.Uk)
		link.ShareID, _ = strconv.Atoi(v.Shareid)
		link.ShortUrl = v.ShortUrl
		sf := &models.SharedFile{}
		sf.Category = category
		sf.Title = v.Title

		err := controllers.SendSharedfile(link, sf)
		if err != nil {
			println(err.Error())
		}

	}

}
func newFetcher() []*Record {
	urlstr := "https://pan.baidu.com/pcloud/feed/getdynamiclist?xxxxx"
	req, _ := http.NewRequest("GET", urlstr, nil)
	req.Header.Set("Cookie", "xxxxx")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	auto := &AutoGenerated{}
	err := json.Unmarshal(body, auto)
	if err != nil {
		return nil
	}
	return auto.Records
}

type AutoGenerated struct {
	Errno      int       `json:"errno"`
	RequestID  int       `json:"request_id"`
	TotalCount int       `json:"total_count"`
	Records    []*Record `json:"records"`
}
type Record struct {
	Shareid  string `json:"shareid"`
	Title    string `json:"title"`
	Uk       int    `json:"uk"`
	ShortUrl string `json:"shorturl"`
}
