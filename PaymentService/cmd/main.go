package main

import (
	"Payment/PaymentService/internal/domain"
	"Payment/PaymentService/internal/kafka"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
	"strconv"
	"time"
)

func main() {

	kafkaProducer, err := kafka.NewProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	kafkaConsumer, err := kafka.NewConsumer([]string{"kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Канал для получения сообщений
	responseChannel := make(chan *sarama.ConsumerMessage)
	go kafkaConsumer.Consume("orders", 0, responseChannel)

	go func() {
		for msg := range responseChannel {
			var order domain.Orders
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("Failed to unmarshal order: %v", err)
				continue
			}

			newStatus := processOrder(order)

			response := domain.OrderResponse{
				OrderID:   strconv.Itoa(order.Id),
				NewStatus: newStatus,
			}

			responseData, err := json.Marshal(response)
			if err != nil {
				log.Printf("Failed to marshal response: %v", err)
				continue
			}

			err = kafkaProducer.ProduceMessage("order_responses", response.OrderID, responseData)
			if err != nil {
				log.Printf("Failed to send response to Kafka: %v", err)
			}
		}
	}()

	select {}
}

func processOrder(order domain.Orders) string {

	randomValue := order.Id % 7
	time.Sleep(time.Duration(randomValue) * time.Second)
	if randomValue < 5 {
		return "Paid"
	} else {
		return "Canseled"
	}
}
