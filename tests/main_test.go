package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylannguyennn/url-shortener/database"
	"github.com/gin-gonic/gin"
)

// Test router
func setupRouter() *gin.Engine {
	_ = database.Connect()
	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		var req struct {
			URL string `json:"url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"short_url": "http://localhost/test123",
			"original":  req.URL,
		})
	})

	return r
}

func TestHealthEndpoint(t *testing.T) {
	router := setupRouter()

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
