package models

import (
	"encoding/json"
	"github.com/rafaeljesus/event-tracker/lib/elastic"
	"time"
)

type Event struct {
	Cid       int             `json:"cid"`
	Name      string          `json:"name"`
	Status    string          `json:"status"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp,omitempty"`
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
