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
	query := buildQuery(q)

	searchResult, err := query.
		Sort("timestamp", true).
		From(0).Size(10).
		Pretty(true).
		Do()

	if err != nil {
		return err, nil
	}

	result := parseResult(searchResult)

	return nil, result
}

func parseResult(searchResult *client.SearchResult) []Event {
	var result []Event

	for _, hit := range searchResult.Hits.Hits {
		var event Event
		json.Unmarshal(*hit.Source, &event)
		result = append(result, event)
	}

	return result
}

func buildQuery(q Query) *client.SearchService {
	index := elastic.Es.Search().Index("events")

	if q.Cid != "" {
		cid := client.NewTermQuery("cid", q.Cid)
		index.Query(cid)
	}

	if q.Name != "" {
		name := client.NewTermQuery("name", q.Name)
		index.Query(name)
	}

	if q.Status != "" {
		status := client.NewTermQuery("status", q.Status)
		index.Query(status)
	}

	return index
}
