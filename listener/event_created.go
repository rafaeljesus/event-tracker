package listener

import (
	"encoding/json"
	"github.com/rafaeljesus/event-tracker/models"
	"log"
)

func EventCreated(message []byte) {
	event := &models.Event{}

	if err := json.Unmarshal(message, &event); err != nil {
		log.Fatalln("Failed to parse message", err)
	}
}
