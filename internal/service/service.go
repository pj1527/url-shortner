package service

import (
	"context"
	"fmt"
	"url-shortener/internal/repository"
	"url-shortener/pkg/utils"
)

type Service struct {
	Repository repository.Repository
}

func NewService(repository repository.Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) GenerateShortKey(ctx context.Context, longURL string) (string, error) {
	id, err := s.Repository.SaveURL(ctx, longURL)
	if err != nil {
		return "", err
	}
	shortKey := utils.ToBase62(id)
	return shortKey, nil
}

func (s *Service) FetchLongURL(ctx context.Context, shortKey string) (string, error) {
	id := utils.ToInteger(shortKey)
	longURL, ok := s.Repository.GetURL(ctx, id)
	if !ok {
		return "", fmt.Errorf("URL not found")
	}
	return longURL, nil
}
