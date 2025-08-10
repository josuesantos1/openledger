package handler

import (
	"log"
	"github.com/josuesantos1/openledger/pkg"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

type ConsumerHandler struct {
	storage *pkg.Storage
	queue   *pkg.RabbitMQ
}

func NewConsumerHandler(storage *pkg.Storage, queue *pkg.RabbitMQ) *ConsumerHandler {
	return &ConsumerHandler{
		storage: storage,
		queue:   queue,
	}
}

func (h *ConsumerHandler) RegisterConsumers() error {
	queues := map[string]func(amqp091.Delivery){
		"openledger.transaction": h.Transaction,
		"openledger.commit":      h.Commit,
	}

	for queueName, handler := range queues {
		q, err := h.queue.QueueDeclare(queueName, false, false, false, false, nil)
		if err != nil {
			return err
		}

		go func(queueName string, handler func(amqp091.Delivery)) {
			msgs, err := h.queue.Channel.Consume(
				q.Name,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				log.Printf("Error while consuming messages from queue %s: %v", q.Name, err)
				return
			}

			for d := range msgs {
				handler(d)
			}
		}(queueName, handler)
	}

	return nil
}

func (h *ConsumerHandler) Transaction(d amqp091.Delivery) {
	log.Printf("Processing transaction message: %s", d.Body)
}

func (h *ConsumerHandler) Commit(d amqp091.Delivery) {
	log.Printf("Processing commit message: %s", d.Body)
}
