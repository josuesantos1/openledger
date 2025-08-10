package handler

import (
	"github.com/josuesantos1/openledger/internal/domain"
	"log"
	"time"
	"net/http"
	"encoding/json"
	"github.com/josuesantos1/openledger/pkg"
)

type AccountHandler struct {
	storage *pkg.Storage
}

func NewAccountHandler(storage *pkg.Storage) *AccountHandler {
	return &AccountHandler{
		storage: storage,
	}
}

func (h *AccountHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/accounts", h.CreateAccount)
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var account domain.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := time.Now().Format("20060102150405.000000000")
	account.ID = id
	account.CreatedAt = time.Now()

	data, err := json.Marshal(account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.storage.Save(id, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Account created: %+v", account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(account); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing Account ID", http.StatusBadRequest)
		return
	}

	data, err := h.storage.Load(id)
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	log.Printf("Account loaded: %s", string(data))

	var Account domain.Account
	if err := json.Unmarshal(data, &Account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Account); err != nil {
		log.Printf("Erro ao codificar resposta: %v", err)
	}
}

