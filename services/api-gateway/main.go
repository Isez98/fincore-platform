package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Println("error writing response:", err)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Println("Listening on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
