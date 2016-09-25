package db

import (
	"gopkg.in/olivere/elastic.v3"
	"log"
)

var Es *elastic.Client

func Connect() {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	_, err = client.CreateIndex("events").Do()
	if err != nil {
		log.Print(err)
	}

	Es = client
}
