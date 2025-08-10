package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/josuesantos1/openledger/internal/domain"
)

type ClientHandler struct {

}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{}

}


func (h *ClientHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/clients", h.CreateClient)
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client domain.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Client created:", client)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}
