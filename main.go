package main

import (
	"go-search-string/ieee"
	"log"
	"os"
)

func main() {
	Key := os.Getenv("Key")

	if Key == "" {
		log.Fatalln("[ERROR] Set env variable Key with Ieee api key")
	}

	api := ieee.IeeeApi{Key: Key}

	res, err := api.SearchMock("(Model based test) AND (web application)")

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Articles[0])
}
