package service

import (
	"fmt"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type Service struct {
	Repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GenerateShortKey(longURL string) (string, error) {
	id := s.Repository.SaveURL(longURL)
	shortKey := utils.ToBase62(id)
	return shortKey, nil
}

func (s *Service) FetchLongURL(shortKey string) (string, error) {
	id := utils.ToInteger(shortKey)
	longURL, ok := s.Repository.GetURL(id)
	if !ok {
		return "", fmt.Errorf("URL not found")
	}
	return longURL, nil
}
