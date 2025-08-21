package repository

import "sync"

type Repository struct {
	mu      *sync.RWMutex
	urls    map[uint64]string
	counter uint64
}

func NewRepository() *Repository {
	return &Repository{
		mu:      &sync.RWMutex{},
		urls:    make(map[uint64]string),
		counter: 10000000,
	}
}

func (s *Repository) SaveURL(longURL string) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	s.urls[s.counter] = longURL
	return s.counter
}

func (s *Repository) GetURL(id uint64) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, ok := s.urls[id]
	return longURL, ok
}
