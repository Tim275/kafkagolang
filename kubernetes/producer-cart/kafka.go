package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

// Order repräsentiert eine Bestellung
type Order struct {
	OrderID    string  `json:"order_id"`
	Customer   string  `json:"customer"`
	Product    string  `json:"product"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

// KafkaProducer definiert den Kafka-Writer
type KafkaProducer struct {
	writer *kafka.Writer
}

// NewKafkaProducer initialisiert einen neuen Kafka-Writer
func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll, // Warten auf alle Acks
		},
	}
}

// Close schließt den Kafka-Writer
func (kp *KafkaProducer) Close() error {
	return kp.writer.Close()
}

// SendOrder sendet eine Bestellung an Kafka
func (kp *KafkaProducer) SendOrder(order Order) error {
	msgBytes, err := json.Marshal(order)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(order.OrderID),
		Value: msgBytes,
	}

	return kp.writer.WriteMessages(context.Background(), msg)
}
