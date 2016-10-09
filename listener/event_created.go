package listener

import (
	"log"
)

func EventCreated(payload string) {
	log.Println("Receiving Event Message", payload)

	return
}
