package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Handler für HTTP-Anfragen zum Senden von Bestellungen
func (kp *KafkaProducer) orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Methode nicht erlaubt", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&order)
	if err != nil {
		http.Error(w, "Ungültige Anfrage", http.StatusBadRequest)
		return
	}

	err = kp.SendOrder(order)
	if err != nil {
		log.Printf("Fehler beim Senden der Bestellung: %v", err)
		http.Error(w, "Fehler beim Senden der Bestellung", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Bestellung %s gesendet", order.OrderID)
}

func main() {
	brokers := []string{os.Getenv("KAFKA_BROKERS")}
	topic := os.Getenv("KAFKA_TOPIC")

	producer := NewKafkaProducer(brokers, topic)
	defer producer.Close()

	http.HandleFunc("/send-order", producer.orderHandler)

	serverPort := "8080" // Standardport
	if port := os.Getenv("PORT"); port != "" {
		serverPort = port
	}

	srv := &http.Server{
		Addr:         ":" + serverPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Producer API startet auf Port %s...", serverPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server fehlgeschlagen: %v", err)
	}
}
