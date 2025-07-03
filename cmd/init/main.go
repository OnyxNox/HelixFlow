package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	for i := range 10 {
		conn, err := net.DialTimeout("tcp", "kafka:9092", 2*time.Second)
		if err == nil {
			conn.Close()

			log.Println("Kafka is accepting connections.")

			break
		}

		log.Println("Waiting for Kafka to be ready...")

		time.Sleep(3 * time.Second)

		if i == 9 {
			log.Fatal("Kafka did not become ready in time.")
		}
	}

	broker := "kafka:9092"
	topics := []kafka.TopicConfig{
		{
			Topic:             "jobs.schedule",
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
		{
			Topic:             "jobs.deadletter",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		{
			Topic:             "status.updates",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := kafka.DialContext(ctx, "tcp", broker)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Kafka: %v", err))
	}
	defer conn.Close()

	for _, topic := range topics {
		fmt.Printf("Creating topic: %s\n", topic.Topic)

		if err := conn.CreateTopics(topic); err != nil {
			fmt.Printf("Failed to create topic %s: %v\n", topic.Topic, err)
		} else {
			fmt.Printf("Topic created: %s\n", topic.Topic)
		}
	}
}
