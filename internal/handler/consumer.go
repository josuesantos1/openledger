package handler

import (
	"log"
	"github.com/josuesantos1/openledger/pkg"
)

type ConsumerHandler struct {
	storage *pkg.Storage
}

func NewConsumerHandler(storage *pkg.Storage) *ConsumerHandler {
	return &ConsumerHandler{
		storage: storage,
	}
}

func (h *ConsumerHandler) RegisterConsumers(queue *pkg.RabbitMQ) error {
	q, err := queue.QueueDeclare("openledger.transaction", false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := queue.Channel.Consume(
		q.Name,    
		"",  
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	var done chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-done
	return nil
}

