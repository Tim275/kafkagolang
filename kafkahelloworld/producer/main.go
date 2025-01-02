package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	
	brokers := []string{"localhost:29092"}
	topic := "TEST_TOPIC"

	err := createTopic(brokers, topic)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Topic created")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	ctx := context.Background()
	err = writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World 4!"),
		})

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Successfully published message to topic")
}