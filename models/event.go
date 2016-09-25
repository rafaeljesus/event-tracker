package models

import (
	"github.com/rafaeljesus/event-tracker/db"
	"time"
)

type Event struct {
	Cid       int       `json:"cid"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (e *Event) Create() error {
	_, err := db.Es.Index().
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
