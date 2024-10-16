package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

type Producer struct {
	sarama.SyncProducer
}

type Consumer struct {
	sarama.Consumer
}

// Функция для создания нового продюсера
func NewProducer(addresses []string) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(addresses, nil)
	if err != nil {
		return nil, err
	}
	return &Producer{SyncProducer: producer}, nil
}

// Функция отправки сообщения в Kafka
func (p *Producer) ProduceMessage(topic string, orderID string, orderData interface{}) error {
	// Преобразование объекта заказа в JSON
	orderDataJSON, err := json.Marshal(orderData)
	if err != nil {
		return err
	}

	// Создание сообщения
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(orderID),
		Value: sarama.ByteEncoder(orderDataJSON),
	}

	// Отправка сообщения в Kafka
	_, _, err = p.SendMessage(msg)
	return err
}

// Функция для создания нового консьюмера
func NewConsumer(addresses []string) (*Consumer, error) {
	consumer, err := sarama.NewConsumer(addresses, nil)
	if err != nil {
		return nil, err
	}
	return &Consumer{Consumer: consumer}, nil
}

// Метод для обработки сообщений из Kafka
func (c *Consumer) Consume(topic string, partition int32, messageChan chan<- *sarama.ConsumerMessage) {
	partConsumer, err := c.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Printf("Failed to consume partition: %v", err)
		return
	}
	defer partConsumer.Close()

	for msg := range partConsumer.Messages() {
		messageChan <- msg
	}
}
