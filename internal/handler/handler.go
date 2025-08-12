package handler

import (
	"fmt"
	"io"
	"net/http"
	"url-shortener/internal/service"
	"url-shortener/internal/store"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		Service: service.NewService(store),
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)
}

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	longURL, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortURL, err := h.Service.GenerateShortURL(scheme, r.Host, string(longURL))
	if err != nil {
		http.Error(w, "Error shortening URL", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", shortURL)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[1:]
	if shortKey == "" {
		fmt.Fprintf(w, "Welcome to the URL Shortener API")
		return
	}
	longURL, err := h.Service.FetchLongURL(shortKey)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
