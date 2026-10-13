package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	repo := &InMemoryHelathRepository{}
	service := NewHealthService(repo)

	response := HealthResponse{
		Status: service.Check(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode health response: %v", err)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	taskRepo := NewInMemoryTaskRepository()
	taskHandler := NewTaskHandler(taskRepo)

	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks", taskHandler.List)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.Delete)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           loggingMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("server is running on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
