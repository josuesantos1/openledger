package pkg

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit interface {
	Connect() error
	Disconnect() error
	QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args *amqp.Table) (amqp.Queue, error)
}

type RabbitMQ struct {
	URI      string
	Username string
	Password string
	Conn     *amqp.Connection
	Channel  *amqp.Channel
}

func NewRabbitMQ(uri string, username string, password string) *RabbitMQ {
	return &RabbitMQ{
		URI:      uri,
		Username: username,
		Password: password,
	}
}

func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial(r.URI)
	if err != nil {
		return err
	}
	
	r.Conn = conn

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	r.Channel = channel

	return nil
}

func (r *RabbitMQ) Disconnect() error {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			return err
		}
	}
	
	if r.Conn != nil {
		if err := r.Conn.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (r *RabbitMQ) QueueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) (amqp.Queue, error) {
	return r.Channel.QueueDeclare(
		name,
		durable,
		autoDelete,
		exclusive,
		noWait,
		args,
	)
}

