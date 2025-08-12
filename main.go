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
	h.RegisterRoutes(http.DefaultServeMux)

	fmt.Println("Server starting on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
