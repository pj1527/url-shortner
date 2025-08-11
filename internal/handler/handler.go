package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"url-shortener/internal/service"
	"url-shortener/internal/store"
)

type Handler struct {
	Store *store.Store
}

func NewHandler(store *store.Store) *Handler {
	return &Handler{
		Store: store,
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

	if _, err := url.ParseRequestURI(string(longURL)); err != nil {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	id := h.Store.SaveURL(string(longURL))
	shortKey := service.ToBase62(id)

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)
	fmt.Fprintf(w, "%s", shortURL)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[1:]
	if shortKey == "" {
		fmt.Fprintf(w, "Welcome to the URL Shortener API")
		return
	}

	id := service.ToInteger(shortKey)
	longURL, ok := h.Store.GetURL(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
