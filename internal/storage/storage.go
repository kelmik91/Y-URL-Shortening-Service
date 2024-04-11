package storage

import "sync"

var StorageURL = Storage{
	storage: make(map[string]string),
	mu:      sync.RWMutex{},
}

type Storage struct {
	storage map[string]string
	mu      sync.RWMutex
}

func (s *Storage) AddURL(id, url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[id] = url
}

func (s *Storage) Get(id string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.storage[id]
}

func New() *Storage {
	return &Storage{
		storage: make(map[string]string),
		mu:      sync.RWMutex{},
	}
}
