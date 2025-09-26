package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker = "localhost:9092"
	topic       = "student.registered"
)

// Event represents the student registration event structure
type Event struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	log.Println("[StudentService] Starting service...")

	// Create a new writer to produce messages
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Simulate a new student registration every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	studentID := 1
	for range ticker.C {
		// Create a new registration event
		event := Event{
			StudentID: "S" + time.Now().Format("20060102150405"),
			Name:      "Student-" + string(rune(studentID)),
			Status:    "registered",
		}
		studentID++

		// Marshal the event into JSON
		eventBytes, err := json.Marshal(event)
		if err != nil {
			log.Printf("[StudentService] Failed to marshal event: %v", err)
			continue
		}

		// Send the event to Kafka
		msg := kafka.Message{
			Key:   []byte(event.StudentID),
			Value: eventBytes,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = writer.WriteMessages(ctx, msg)
		cancel()

		if err != nil {
			log.Printf("[StudentService] Failed to send message: %v", err)
		} else {
			log.Printf("[StudentService] Sent event: %s, student_id: %s", topic, event.StudentID)
		}
	}
}