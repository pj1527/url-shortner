package store

import "sync"

type Store struct {
	mu       *sync.RWMutex
	urlStore map[uint64]string
	counter  uint64
}

func NewStore() *Store {
	return &Store{
		mu:       &sync.RWMutex{},
		urlStore: make(map[uint64]string),
		counter:  10000000,
	}
}

func (s *Store) SaveURL(longURL string) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter++
	s.urlStore[s.counter] = longURL
	return s.counter
}

func (s *Store) GetURL(id uint64) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	longURL, ok := s.urlStore[id]
	return longURL, ok
}
