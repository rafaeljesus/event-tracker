package kafka

import (
	"github.com/Shopify/sarama"
)

var brokers = []string{"localhost:9092"}
var Producer sarama.SyncProducer
var Consumer sarama.Consumer

func Connect() error {
	consumer, _ := sarama.NewConsumer(brokers, nil)
	Consumer = consumer

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, _ := sarama.NewSyncProducer(brokers, config)
	Producer = producer

	return nil
}

func Enqueue(topic string, msg []byte) error {
	message := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}

	_, _, err := Producer.SendMessage(message)
	return err
}

func FromQueue(topic string, fn func(string)) error {
	pc, _ := Consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	go func(pc sarama.PartitionConsumer) {
		for message := range pc.Messages() {
			fn(string(message.Value))
		}
	}(pc)

	return nil
}
