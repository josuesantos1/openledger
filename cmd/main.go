
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/josuesantos1/openledger/pkg"
	"github.com/josuesantos1/openledger/internal/handler"
)

func main() {
	server := pkg.NewHTTPServer(":8081")
	mux := server.Server()

	clientHandler := handler.NewClientHandler()
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Stop(); err != nil {
		log.Printf("Falha ao encerrar a aplicação: %v", err)
	}
}