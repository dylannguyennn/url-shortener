package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylannguyennn/url-shortener/database"
)

func TestHealthEndpoint(t *testing.T) {
	_ = database.Connect()
	router := main.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 but got %d", w.Code)
	}
}

func TestShortenEndpoint(t *testing.T) {
	router := setupRouter()

	body := map[string]string{"url": "https://github.com"}
	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected 200 but got %d", w.Code)
	}
}
