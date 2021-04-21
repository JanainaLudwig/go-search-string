package ieee

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const baseUri = "http://ieeexploreapi.ieee.org/api/v1/search/articles"

type IeeeApi struct {
	Key string
}

type IeeeResponse struct {
	TotalRecords int `json:"total_records"`
	TotalSearched int `json:"total_searched"`
	Articles []IeeeArticle `json:"articles"`
}
type IeeeArticle struct {
	Abstract string `json:"abstract"`
	Doi string `json:"doi"`
	HtmlUrl string `json:"html_url"`
	Title string `json:"title"`
}

func (i *IeeeApi) Search(search string) (response *IeeeResponse, err error) {
	//log.Println("[SEARCHING]", baseUri + "?format=json&apikey=" + i.Key + "&querytext=" + url.QueryEscape(search))
	res, err := http.Get(baseUri + "?format=json&apikey=" + i.Key + "&querytext=" + url.QueryEscape(search))

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (i *IeeeApi) SearchMock(search string) (response *IeeeResponse, err error) {
	err = json.Unmarshal([]byte(Example), &response)

	return response, err
}
