package repository

import (
	"errors"
	"sync"
)

var ErrURLNotFound = errors.New("URL not found")

type URLRepository interface {
	Save(code string, longUrl string) error
	Find(code string) (string, error)
}

type inMemoryURLRepository struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewInMemoryURLRepository() URLRepository {
	return &inMemoryURLRepository{
		store: make(map[string]string),
	}
}

func (r *inMemoryURLRepository) Save(code string, longUrl string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[code] = longUrl
	return nil
}

func (r *inMemoryURLRepository) Find(code string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	longUrl, exists := r.store[code]
	if !exists {
		return "", ErrURLNotFound
	}
	return longUrl, nil
}
