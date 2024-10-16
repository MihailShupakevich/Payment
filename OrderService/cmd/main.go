package main

import (
	"Payment/OrderService/internal/db"
	"Payment/OrderService/internal/domain"
	"Payment/OrderService/internal/handler"
	"Payment/OrderService/internal/kafka"
	"Payment/OrderService/internal/repository"
	"Payment/OrderService/internal/usecase"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	databaseConnect, err := db.Db()
	if err != nil {
		panic(err)
	}

	orderRepo := repository.New(databaseConnect)
	orderUsecase := usecase.New(orderRepo)

	kafkaProducer, err := kafka.NewProducer([]string{"kafka:9092"})
	if err != nil {
		panic(err)
	}
	defer kafkaProducer.Close()

	kafkaConsumer, err := kafka.NewConsumer([]string{"kafka:9092"})
	if err != nil {
		panic(err)
	}

	// Канал для получения сообщений
	responseChannel := make(chan *sarama.ConsumerMessage)
	go kafkaConsumer.Consume("order_responses", 0, responseChannel) // Подписка на топик ответов

	orderHandler := handler.New(orderUsecase, kafkaProducer)

	// Горутина для обработки ответов
	go func() {
		for msg := range responseChannel {
			var response domain.OrderResponse
			if err := json.Unmarshal(msg.Value, &response); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			// Обновить статус заказа на основе ответа
			orderID, err := strconv.Atoi(response.OrderID)
			if err != nil {
				log.Printf("Invalid Order ID: %v", err)
				continue
			}

			_, err = orderHandler.UpdateOrder(orderID, response.NewStatus)
			if err != nil {
				log.Printf("Failed to update order status: %v", err)
			} else {
				log.Printf("Order %d updated to status: %s", orderID, response.NewStatus)
			}
		}
	}()

	router := gin.Default()

	orderGroup := router.Group("/orders")
	orderHandler.SetupRoutes(orderGroup)
	router.Run(":8080")
}
