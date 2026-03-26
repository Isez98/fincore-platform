package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Println("Server healthy")
	})

	log.Println("Listending on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
