package repository

import (
	"context"
	"sync"
)

type inMemoryRepository struct {
	mu      *sync.RWMutex
	urls    map[uint64]string
	counter uint64
}

func NewInMemoryRepository() (Repository, error) {
	return &inMemoryRepository{
		mu:      &sync.RWMutex{},
		urls:    make(map[uint64]string),
		counter: 10000000,
	}, nil
}

func (s *inMemoryRepository) SaveURL(ctx context.Context, longURL string) (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	s.urls[s.counter] = longURL
	return s.counter, nil
}

func (s *inMemoryRepository) GetURL(ctx context.Context, id uint64) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, ok := s.urls[id]
	return longURL, ok
}
