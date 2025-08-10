package handler

import (
	"encoding/json"
	"github.com/josuesantos1/openledger/internal/domain"
	"github.com/josuesantos1/openledger/pkg"
	"log"
	"net/http"
	"time"
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
	mux.HandleFunc("GET /v1/clients", h.GetClient)
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var client domain.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := time.Now().Format("20060102150405.000000000")
	client.ID = id
	client.CreatedAt = time.Now()

	data, err := json.Marshal(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.storage.Save(id, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Client created: %+v", client)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(client); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err)
	}
}

func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing client ID", http.StatusBadRequest)
		return
	}

	data, err := h.storage.Load(id)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	log.Printf("Client loaded: %s", string(data))

	var client domain.Client
	if err := json.Unmarshal(data, &client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(client); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err)
	}
}
