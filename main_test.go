package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	expected := "{\"status\":\"healthy\"}\n"

	if rec.Body.String() != expected {
		t.Fatalf("expected %s, got %s", expected, rec.Body.String())

	}
}
