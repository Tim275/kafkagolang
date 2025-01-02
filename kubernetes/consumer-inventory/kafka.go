package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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

// KafkaConsumer definiert den Kafka-Reader
type KafkaConsumer struct {
	reader  *kafka.Reader
	topic   string
	groupID string
}

// NewKafkaConsumer initialisiert einen neuen Kafka-Reader
func NewKafkaConsumer(brokers []string, topic, groupID string) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumer{
		reader:  reader,
		topic:   topic,
		groupID: groupID,
	}
}

// Close schließt den Kafka-Reader
func (kc *KafkaConsumer) Close() error {
	return kc.reader.Close()
}

// ConsumeMessages liest und verarbeitet Nachrichten von Kafka
func (kc *KafkaConsumer) ConsumeMessages() {
	fmt.Println("Inventory Service startet und konsumiert Bestellungen...")
	for {
		msg, err := kc.reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Fehler beim Lesen der Nachricht: %s", err)
			break // Verbindung neu aufbauen
		}

		var order Order
		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Printf("Fehler beim Deserialisieren der Bestellung: %s", err)
			continue
		}

		fmt.Printf("Bestellung erhalten: %+v\n", order)
		fmt.Printf("Artikel '%s' in Menge %d reserviert für Bestellung %s\n", order.Product, order.Quantity, order.OrderID)
	}
}
