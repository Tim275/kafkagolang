package main

import (
	"log"
	"time"
)

func main() {
	// Kafka-Konfiguration
	brokers := []string{"kafka:9092"}
	topic := "ORDER_TOPIC"
	groupID := "shippingGroup"

	for {
		// Kafka Consumer initialisieren
		consumer := NewKafkaConsumer(brokers, topic, groupID)

		// Nachrichten konsumieren
		consumer.ConsumeMessages()

		// Consumer schließen
		err := consumer.Close()
		if err != nil {
			log.Printf("Fehler beim Schließen des Consumers: %v", err)
		}

		log.Println("Versuche erneut, eine Verbindung zu Kafka herzustellen...")
		time.Sleep(5 * time.Second)
	}
}
