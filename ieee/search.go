package ieee

import (
	"log"
	"net/http"
	"net/url"
)

const baseUri = "http://ieeexploreapi.ieee.org/api/v1/search/articles"

type IeeeApi struct {
	Key string
}

func (i *IeeeApi) Search(search string) (*http.Response, error) {
	log.Println(baseUri + "?format=json&apikey=" + i.Key + "&querytext=" + url.QueryEscape(search))
	return http.Get(baseUri + "?format=json&apikey=" + i.Key + "&querytext=" + url.QueryEscape(search))
}
