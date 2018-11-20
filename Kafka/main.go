package main

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produceMessage(topicName string) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		log.Panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topicName,
			Partition: kafka.PartitionAny},
		Value: []byte("my message"),
	}, nil)

	if err != nil {
		log.Fatal("Error on producer: ", err)
	}
	log.Println("Message published")
}

func consumeMessage(topicName string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"group.id":           "myGroup",
		"session.timeout.ms": 6000,
		//"auto.offset.reset":  "earliest",
	})

	if err != nil {
		log.Panic(err)
	}
	c.SubscribeTopics([]string{topicName}, nil)

	log.Println("Consumer started")
	for {
		ev := c.Poll(100)
		if ev != nil {
			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
			default:
				//log.Printf("Ignored %v\n", e)
			}
		}
	}
	log.Println("About to close consumer")
	c.Close()
}

func main() {
	topicName := "test"

	go consumeMessage(topicName)
	produceMessage(topicName)
	time.Sleep(10 * time.Second)
}
