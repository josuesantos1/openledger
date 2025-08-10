package handler

import (
	"net/http"
	"encoding/json"
	"fmt"
	"time"
	"github.com/josuesantos1/openledger/internal/domain"
	"github.com/josuesantos1/openledger/pkg"
)

type ClientHandler struct {
	storage *pkg.Storage
}

func NewClientHandler(storage *pkg.Storage) *ClientHandler {
	return &ClientHandler{
		storage: storage,
	}
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

	id := fmt.Sprintf("%d", time.Now().UnixNano())

	client.ID = id
	client.CreatedAt = time.Now()

	if err := h.storage.Save(id, &client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Client created:", client)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}
