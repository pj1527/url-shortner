package main

import (
	"fmt"
	"net/http"
	"url-shortener/internal/handler"
	"url-shortener/internal/store"
)

func main() {
	urlStore := store.NewStore()

	h := handler.NewHandler(urlStore)

	http.HandleFunc("/shorten", h.Shorten)
	http.HandleFunc("/", h.Redirect)

	fmt.Println("Server starting on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
