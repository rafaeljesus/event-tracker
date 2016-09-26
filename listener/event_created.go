package listener

import (
	"github.com/Shopify/sarama"
	"log"
)

func EventCreated(message *sarama.ConsumerMessage) {
	log.Print("Receiving Event Message")
}
