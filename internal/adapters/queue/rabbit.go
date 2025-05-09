package queue

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(queueName string) (*RabbitMQ, error) {
	host := os.Getenv("RABBITMQ_HOST")
	if host == "" {
		host = "rabbitmq"
	}

	port := os.Getenv("RABBITMQ_PORT")
	if port == "" {
		port = "5672"
	}

	user := os.Getenv("RABBITMQ_USER")
	if user == "" {
		user = "guest"
	}

	password := os.Getenv("RABBITMQ_PASSWORD")
	if password == "" {
		password = "guest"
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	var conn *amqp.Connection
	var err error
	maxRetries := 5
	retryDelay := time.Second * 5

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		log.Printf("Tentativa %d de %d de conectar ao RabbitMQ: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao RabbitMQ após %d tentativas: %v", maxRetries, err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName, true, false, false, false, nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{conn, ch, q}, nil
}

func (r *RabbitMQ) Publish(event []byte) error {
	return r.channel.Publish(
		"", r.queue.Name, false, false,
		amqp.Publishing{ContentType: "application/json", Body: event},
	)
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}
