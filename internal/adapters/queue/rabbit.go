package queue

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		return nil, err
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
