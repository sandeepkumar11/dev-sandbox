package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/repository"
	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/service"
)

type ShortenHandler struct {
	service service.ShortenService
}

func NewShortenHandler(s service.ShortenService) *ShortenHandler {
	return &ShortenHandler{
		service: s,
	}
}

func (h *ShortenHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.Shorten(req.URL)
	if err != nil {
		if err == service.ErrInvalidURL {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{"short_url": shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *ShortenHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")

	longURL, err := h.service.Resolve(code)
	if err != nil {
		if err == repository.ErrURLNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
