package service

import (
	"fmt"
	"net/url"
	"url-shortener/internal/store"
	"url-shortener/pkg/utils"
)

type Service struct {
	Store *store.Store
}

func NewService(store *store.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GenerateShortURL(httpScheme string, domain string, longURL string) (string, error) {
	if _, err := url.ParseRequestURI(string(longURL)); err != nil {
		return "", fmt.Errorf("invalid URL format: %v", err)
	}
	id := s.Store.SaveURL(string(longURL))
	shortKey := utils.ToBase62(id)
	shortURL := fmt.Sprintf("%s://%s/%s", httpScheme, domain, shortKey)
	return shortURL, nil
}

func (s *Service) FetchLongURL(shortKey string) (string, error) {
	id := utils.ToInteger(shortKey)
	longURL, ok := s.Store.GetURL(id)
	if !ok {
		return "", fmt.Errorf("URL not found")
	}
	return longURL, nil
}
