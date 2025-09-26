package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaBroker            = "localhost:9092"
	registeredTopic        = "student.registered"
	validatedTopic         = "student.registration_validated"
	failedTopic            = "student.registration_failed"
)

// Event represents the event structure
type Event struct {
	StudentID string `json:"student_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

func main() {
	log.Println("[FinanceService] Starting service...")

	// Create a reader to consume messages from 'student.registered' topic
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   registeredTopic,
		GroupID: "finance-service-group", // Unique GroupID for this service
	})
	defer reader.Close()

	// Create writers to produce messages to different topics
	validatedWriter := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    validatedTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer validatedWriter.Close()

	failedWriter := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    failedTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer failedWriter.Close()

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	for {
		// Read a message from the topic
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		msg, err := reader.FetchMessage(ctx)
		cancel()

		if err != nil {
			log.Printf("[FinanceService] Error fetching message: %v", err)
			continue
		}

		log.Printf("[FinanceService] Received event: %s, student_id: %s", registeredTopic, string(msg.Key))

		var event Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("[FinanceService] Failed to unmarshal event: %v", err)
			continue
		}

		// Simulate payment validation (50% chance of success)
		if rand.Intn(2) == 1 {
			// Payment success
			log.Printf("[FinanceService] Payment validated for student_id: %s", event.StudentID)
			event.Status = "payment_validated"
			eventBytes, _ := json.Marshal(event)

			ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
			err = validatedWriter.WriteMessages(ctx, kafka.Message{Key: msg.Key, Value: eventBytes})
			cancel()

			if err != nil {
				log.Printf("[FinanceService] Failed to send student.registration_validated event: %v", err)
			} else {
				log.Printf("[FinanceService] Sent event: %s", validatedTopic)
			}

		} else {
			// Payment failure
			log.Printf("[FinanceService] Payment failed for student_id: %s", event.StudentID)
			event.Status = "payment_failed"
			eventBytes, _ := json.Marshal(event)

			ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
			err = failedWriter.WriteMessages(ctx, kafka.Message{Key: msg.Key, Value: eventBytes})
			cancel()

			if err != nil {
				log.Printf("[FinanceService] Failed to send student.registration_failed event: %v", err)
			} else {
				log.Printf("[FinanceService] Sent event: %s", failedTopic)
			}
		}

		// Commit the message to indicate it has been processed
		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("[FinanceService] Failed to commit message: %v", err)
		}
	}
}