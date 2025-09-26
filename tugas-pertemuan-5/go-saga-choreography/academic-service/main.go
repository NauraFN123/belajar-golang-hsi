package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker      = "localhost:9092"
	validatedTopic   = "student.registration_validated"
	initializedTopic = "student.academic_initialized"
)

// Event represents the event structure
type Event struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	log.Println("[AcademicService] Starting service...")

	// Create a reader to consume messages from 'student.registration_validated' topic
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   validatedTopic,
		GroupID: "academic-service-group", // Unique GroupID
	})
	defer reader.Close()

	// Create a writer to produce messages to 'student.academic_initialized' topic
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    initializedTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	for {
		// Read a message from the topic
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		msg, err := reader.FetchMessage(ctx)
		cancel()

		if err != nil {
			log.Printf("[AcademicService] Error fetching message: %v", err)
			continue
		}

		log.Printf("[AcademicService] Received event: %s, student_id: %s", validatedTopic, string(msg.Key))

		var event Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("[AcademicService] Failed to unmarshal event: %v", err)
			continue
		}

		// Process academic setup
		log.Printf("[AcademicService] Initializing academic data for student_id: %s", event.StudentID)
		time.Sleep(1 * time.Second) // Simulate processing time

		// Send event to mark the end of the saga for this process
		event.Status = "academic_initialized"
		eventBytes, _ := json.Marshal(event)

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		err = writer.WriteMessages(ctx, kafka.Message{Key: msg.Key, Value: eventBytes})
		cancel()

		if err != nil {
			log.Printf("[AcademicService] Failed to send student.academic_initialized event: %v", err)
		} else {
			log.Printf("[AcademicService] Sent event: %s", initializedTopic)
		}

		// Commit the message
		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("[AcademicService] Failed to commit message: %v", err)
		}
	}
}