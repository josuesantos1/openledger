
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/josuesantos1/openledger/pkg"
)

func main() {
	server := pkg.NewHTTPServer(":8080")

	go func() {
		if err := server.Start(); err != nil {
			log.Printf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Stop(); err != nil {
		log.Printf("Falha ao encerrar a aplicação: %v", err)
	}
}