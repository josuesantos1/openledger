package pkg

import (
	"net/http"
	"time"
)

type HTTPServer interface {
	Start() error
	Stop() error
}

type HttpServer struct {
	server *http.Server
}

func NewHTTPServer(addr string) HTTPServer {
	mux := http.NewServeMux()

	return &HttpServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (h *HttpServer) Start() error {
	if err := h.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (h *HttpServer) Stop() error {
	if err := h.server.Close(); err != nil {
		return err
	}
	return nil
}
