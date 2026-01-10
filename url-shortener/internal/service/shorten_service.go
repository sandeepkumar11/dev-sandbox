package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"

	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/repository"
)

var ErrInvalidURL = errors.New("Invalid URL")

type ShortenService interface {
	Shorten(longUrl string) (string, error)
	Resolve(code string) (string, error)
}

type shortenService struct {
	repo    repository.URLRepository
	baseURL string
}

func NewShortenService(repo repository.URLRepository, baseURL string) ShortenService {
	return &shortenService{
		repo:    repo,
		baseURL: baseURL,
	}
}

func (s *shortenService) Shorten(longUrl string) (string, error) {
	if !isValidURL(longUrl) {
		return "", ErrInvalidURL
	}

	code, err := generateCode()
	if err != nil {
		return "", err
	}

	if err := s.repo.Save(code, longUrl); err != nil {
		return "", err
	}

	shortUrl := s.baseURL + "/" + code
	return shortUrl, nil
}

func (s *shortenService) Resolve(code string) (string, error) {
	longUrl, err := s.repo.Find(code)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}

func generateCode() (string, error) {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func isValidURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	return err == nil && u.Scheme != "" && u.Host != ""
}
