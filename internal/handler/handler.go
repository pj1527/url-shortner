package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/service"
	"url-shortener/pkg/utils"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url" binding:"required,url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/shorten", h.Shorten)
	}
	router.GET("/:shortKey", h.Redirect)
}

func (h *Handler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	shortKey, err := h.Service.GenerateShortKey(c.Request.Context(), req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error"})
		return
	}
	shortURL := utils.GenerateURL("http", "localhost:8080", shortKey)
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: shortURL})
}

func (h *Handler) Redirect(c *gin.Context) {
	shortKey := c.Param("shortKey")
	longURL, err := h.Service.FetchLongURL(c.Request.Context(), shortKey)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	c.Redirect(http.StatusFound, longURL)
}
