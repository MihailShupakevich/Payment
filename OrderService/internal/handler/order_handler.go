package handler

import (
	"Payment/OrderService/internal/domain"
	"Payment/OrderService/internal/kafka"
	"Payment/OrderService/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	uc       *usecase.OrderUsecase
	producer *kafka.Producer
	consumer *kafka.Consumer
}

type OrderHandlerInterface interface {
	PostOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
}

func New(usecase *usecase.OrderUsecase, producer *kafka.Producer, consumer *kafka.Consumer) *OrderHandler {
	return &OrderHandler{
		uc:       usecase,
		producer: producer,
		consumer: consumer,
	}
}

func (o *OrderHandler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/post", o.PostOrder)
}

func (o *OrderHandler) PostOrder(c *gin.Context) {
	order := new(domain.Orders)
	c.BindJSON(&order)
	fmt.Println(order)
	newOrder, err := o.uc.PostOrder(*order)
	orderId := strconv.Itoa(newOrder.Id)

	// Отправка сообщения о созданном заказе в Kafka
	if err = o.producer.ProduceMessage("orders", orderId, newOrder); err != nil {
		log.Printf("Failed to send order to Kafka: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send order to Kafka"})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func (o *OrderHandler) UpdateOrder(orderId int, NewStatus string) (domain.Orders, error) {
	fmt.Println("Update Order Ha1")
	response, err := o.uc.UpdateOrder(orderId, NewStatus)
	if err != nil {
		log.Printf("Failed to get response from Kafka: %v", err)
	}
	return response, err
}
