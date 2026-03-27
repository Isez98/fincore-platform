package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Account struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := []Account{
		{
			ID:       "acc_001",
			Name:     "Checking Account",
			Balance:  1250.75,
			Currency: "USD",
		},
		{
			ID:       "acc_002",
			Name:     "Savings Account",
			Balance:  5400.00,
			Currency: "USD",
		},
	}

	writeJSON(w, http.StatusOK, accounts)
}

func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("error writing JSON response:", err)
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/health", healthHandler)
	r.Get("/accounts", accountsHandler)

	log.Println("Listening on port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
