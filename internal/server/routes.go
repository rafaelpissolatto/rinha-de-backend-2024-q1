package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", s.healthHandler)
	r.Post("/clientes/{id}/transacoes", s.createTransactionHandler)
	r.Get("/clientes/{id}/extrato", s.getExtractHandler)

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(map[string]string{"message": "It's healthy"})
	_, _ = w.Write(jsonResp)
}

func (s *Server) createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(map[string]string{"message": "Transaction created"})
	_, _ = w.Write(jsonResp)

}

func (s *Server) GetAllTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// jsonResp, _ := json.Marshal(map[string]string{"message": "All transactions retrieved"})
	// _, _ = w.Write(jsonResp)

	// get all transactions from database
	// return transactions

	transactions, err := s.transactionsRepo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(transactions)
	_, _ = w.Write(jsonResp)
}

func (s *Server) getExtractHandler(w http.ResponseWriter, r *http.Request) {
	// jsonResp, _ := json.Marshal(map[string]string{"message": "Extract retrieved"})
	// _, _ = w.Write(jsonResp)

	// get extract from database
	// get customer id from url
	// get all transactions from customer
	// return transactions

	customerID := chi.URLParam(r, "id")

	// get transactions from database
	transactions, err := s.transactionsRepo.GetByCustomerID(customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, _ := json.Marshal(transactions)
	_, _ = w.Write(jsonResp)
}
