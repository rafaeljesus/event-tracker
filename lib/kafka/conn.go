package kafka

import (
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)

var brokers = []string{"docker:9201"}
var Producer sarama.AsyncProducer
var Consumer sarama.Consumer

func Connect() error {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Println("Failed to create consumer: ", err)
		return err
	}

	Consumer = consumer

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		log.Println("Failed to create producer: ", err)
		return err
	}

	Producer = producer

	return nil
}

func Enqueue(topic string, msg []byte) error {
	defer func() {
		if err := Producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	var producer_err *sarama.ProducerError

ProducerLoop:
	for {
		select {
		case Producer.Input() <- message:
			producer_err = nil
		case err := <-Producer.Errors():
			log.Println("Failed to produce message", err)
			producer_err = err
		case <-signals:
			break ProducerLoop
		}
	}

	return producer_err
}

func FromQueue(topic string, fn func(*sarama.ConsumerMessage)) error {
	partitions, err := Consumer.Partitions(topic)
	if err != nil {
		log.Println("Error retrieving partitions ", err)
		return err
	}

	initialOffset := sarama.OffsetOldest

	for _, partition := range partitions {
		pc, _ := Consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fn(message)
			}
		}(pc)
	}
	return nil
}
