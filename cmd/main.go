package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/josuesantos1/openledger/internal/handler"
	"github.com/josuesantos1/openledger/pkg"
)

func main() {
	ctx := context.Background()

	server := pkg.NewHTTPServer(":8081")
	mux := server.Server()

	storage := pkg.NewStorage()
	if err := storage.Start(); err != nil {
		log.Printf("Error while starting storage: %v", err)
		return
	}
	defer storage.Close()

	graph := pkg.NewGraph("neo4j://localhost:7687", "neo4j", "password")

	if err := graph.Connect(); err != nil {
		log.Printf("Error while starting graph storage: %v", err)
		return
	}
	defer graph.Close(ctx)

	clientHandler := handler.NewClientHandler(storage)
	clientHandler.RegisterRoutes(mux)

	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Error while starting server: %v", err)
		}
	}()

	rabbit := pkg.NewRabbitMQ("amqp://guest:guest@localhost:5672/", "guest", "guest")
	if err := rabbit.Connect(); err != nil {
		log.Printf("Error while connecting to RabbitMQ: %v", err)
		return
	}
	defer rabbit.Disconnect()

	consumerHandler := handler.NewConsumerHandler(storage, rabbit)
	if err := consumerHandler.RegisterConsumers(); err != nil {
		log.Printf("Error while registering consumers: %v", err)
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Stop(); err != nil {
		log.Printf("Falha ao encerrar a aplicação: %v", err)
	}
}
