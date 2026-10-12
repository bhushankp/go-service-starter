package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status: "healthy",
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode health response: %v", err)
	}

}

func main() {

	http.HandleFunc("/health", healthHandler)

	log.Println("server is running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
