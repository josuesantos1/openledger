package pkg

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit interface {
	Connect() error
	Disconnect() error
}

type RabbitMQ struct {
	URI      string
	Username string
	Password string
	Conn     *amqp.Connection
	Channel  *amqp.Channel
}

func NewRabbitMQ(uri string, username string, password string) Rabbit {
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
