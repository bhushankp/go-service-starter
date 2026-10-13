package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestTaskMux() *http.ServeMux {
	mux := http.NewServeMux()
	repo := NewInMemoryTaskRepository()
	handler := NewTaskHandler(repo)

	mux.HandleFunc("POST /tasks", handler.Create)
	mux.HandleFunc("GET /tasks", handler.List)
	mux.HandleFunc("GET /tasks/{id}", handler.GetByID)
	mux.HandleFunc("PUT /tasks/{id}", handler.Update)
	mux.HandleFunc("DELETE /tasks/{id}", handler.Delete)

	return mux
}

func TestTaskCRUD(t *testing.T) {
	mux := newTestTaskMux()

	// 1. Create a task.
	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(`{"title":"Learn Go"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("Create: expected 201, got %d", rec.Code)
	}

	var created Task
	if err := json.NewDecoder(rec.Body).Decode(&created); err != nil {
		t.Fatal(err)
	}

	if created.ID != 1 || created.Title != "Learn Go" {
		t.Fatalf("unexpected created task: %+v", created)
	}

	// 2. Read the task.
	req = httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Get: expected 200, got %d", rec.Code)
	}

	// 3. Update the task.
	req = httptest.NewRequest(
		http.MethodPut,
		"/tasks/1",
		strings.NewReader(
			`{"title":"Master Go","status":"in-progress"}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")

	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Update: expected 200, got %d", rec.Code)
	}

	// 4. Delete the task.
	req = httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("Delete: expected 204, got %d", rec.Code)
	}

	// 5. Confirm the task no longer exists.
	req = httptest.NewRequest(http.MethodGet, "/tasks/1", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("Deleted task: expected 404, got %d", rec.Code)
	}
}

func TestCreateTaskValidation(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{
			name: "invalid JSON",
			body: `{"title":`,
		},
		{
			name: "missing title",
			body: `{"title":" "}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mux := newTestTaskMux()

			req := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				strings.NewReader(tc.body),
			)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Fatalf(
					"expected 400 Bad Request, got %d",
					rec.Code,
				)
			}
		})
	}
}

func TestInvalidTaskID(t *testing.T) {
	mux := newTestTaskMux()

	req := httptest.NewRequest(http.MethodGet, "/tasks/abc", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}
