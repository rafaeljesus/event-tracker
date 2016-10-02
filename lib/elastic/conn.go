package elastic

import (
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
)

var Es *elastic.Client

func Connect() {
	url := elastic.SetURL(os.Getenv("ELASTIC_SEARCH_URL"))
	sniff := elastic.SetSniff(false)
	client, err := elastic.NewClient(sniff, url)
	if err != nil {
		panic(err)
	}

	_, err = client.CreateIndex("events").Do()
	if err != nil {
		log.Print(err)
	}

	Es = client
}
