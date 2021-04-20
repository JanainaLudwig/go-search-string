package main

import (
	"go-search-string/ieee"
	"io/ioutil"
	"log"
)

func main() {
	api := ieee.IeeeApi{Key: Key}

	res, err := api.Search("(Model based test) AND (web application)")

	if err != nil {
		log.Println(err)
	}

	log.Println(res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)
	log.Println(bodyString)
}
