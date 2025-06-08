package handler

import (
	"net/http"

	"github.com/mohit83k/url-shortner/service"

	"github.com/mohit83k/url-shortner/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Shortener *service.URLShortener
}

func NewHandler(s *service.URLShortener) *Handler {
	return &Handler{Shortener: s}
}

func (h *Handler) ShortenURL(c *gin.Context) {
	var req model.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	short := h.Shortener.Shorten(req.URL)
	c.JSON(http.StatusOK, model.URLResponse{ShortURL: short})
}

func (h *Handler) Redirect(c *gin.Context) {
	short := c.Param("short")
	long, ok := h.Shortener.Resolve(short)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusFound, long)
}

func (h *Handler) Metrics(c *gin.Context) {
	c.JSON(http.StatusOK, h.Shortener.TopDomains())
}
