package main

import (
	"github.com/mohit83k/url-shortner/handler"
	"github.com/mohit83k/url-shortner/service"
	"github.com/mohit83k/url-shortner/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	store := store.NewURLStore()
	shortener := service.NewURLShortener(store)
	h := handler.NewHandler(shortener)

	r.POST("/shorten", h.ShortenURL)
	r.GET("/metrics/top-domains", h.Metrics)
	r.GET("/:short", h.Redirect)

	return r
}
