package models

import (
	"encoding/json"
	"github.com/rafaeljesus/event-tracker/lib/elastic"
	client "gopkg.in/olivere/elastic.v3"
	"time"
)

type Event struct {
	Cid       int              `json:"cid"`
	Name      string           `json:"name"`
	Status    string           `json:"status"`
	Payload   *json.RawMessage `json:"payload"`
	Timestamp time.Time        `json:"timestamp,omitempty"`
}

func (e *Event) Create() error {
	_, err := elastic.Es.Index().
		Index("events").
		Type("event").
		BodyJson(e).
		Refresh(true).
		Do()

	if err != nil {
		return err
	}

	return nil
}

func Search(q Query) (error, []Event) {
	query := client.NewTermQuery("name", q.Name)
	searchResult, err := elastic.Es.Search().
		Index("events").
		Query(query).
		Sort("timestamp", false).
		From(0).Size(10).
		Do()

	if err != nil {
		return err, nil
	}

	var result []Event

	for _, hit := range searchResult.Hits.Hits {
		var event Event
		json.Unmarshal(*hit.Source, &event)
		result = append(result, event)
	}

	return nil, result
}
