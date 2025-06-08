package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohit83k/url-shortner/service"
	"github.com/mohit83k/url-shortner/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	s := service.NewURLShortener(store.NewURLStore())
	h := NewHandler(s)

	r.POST("/shorten", h.ShortenURL)
	r.GET("/:short", h.Redirect)
	r.GET("/metrics/top-domains", h.Metrics)

	return r
}

func TestShortenAndRedirect(t *testing.T) {
	router := setupTestRouter()

	body := map[string]string{"url": "https://golang.org"}
	jsonBody, _ := json.Marshal(body)

	// POST /shorten
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	var result map[string]string
	_ = json.Unmarshal(resp.Body.Bytes(), &result)
	short := result["short_url"]
	assert.NotEmpty(t, short)

	// GET /:short
	req2 := httptest.NewRequest(http.MethodGet, "/"+short, nil)
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)

	assert.Equal(t, 302, resp2.Code)
	assert.Equal(t, "https://golang.org", resp2.Header().Get("Location"))
}
