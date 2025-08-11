package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func handleShorten(w http.ResponseWriter, r *http.Request) {
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

	mu.Lock()
	counter++
	shortKey := toBase62(counter)
	urlStore[shortKey] = string(longURL)
	mu.Unlock()

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)
	fmt.Fprintf(w, shortURL)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[1:]
	if shortKey == "" {
		fmt.Fprintf(w, "Welcome to the URL Shortener API")
		return
	}

	mu.RLock()
	longURL, ok := urlStore[shortKey]
	mu.RUnlock()

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Server starting on port 8080...")
	_ = http.ListenAndServe(":8080", nil)
}
