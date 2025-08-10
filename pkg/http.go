package pkg

import (
	"net/http"
	"time"
)

type HTTPServer interface {
	Start() error
	Stop() error
	Server() *http.ServeMux
}

type HttpServer struct {
	server *http.Server
	addr   string
	Mux    *http.ServeMux
}

func NewHTTPServer(addr string) HTTPServer {
	return &HttpServer{
		Mux:  http.NewServeMux(),
		addr: addr,
	}
}

func (h *HttpServer) Start() error {

	h.server = &http.Server{
		Addr:         h.addr,
		Handler:      h.Mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

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

func (h *HttpServer) Server() *http.ServeMux {
	return h.Mux
}
