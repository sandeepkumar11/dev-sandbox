package app

import (
	"net/http"

	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/config"
	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/handler"
	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/repository"
	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/service"
)

func SetupRouter(cfg *config.Config) http.Handler {
	repo := repository.NewInMemoryURLRepository()
	svc := service.NewShortenService(repo, cfg.BaseURL)
	h := handler.NewShortenHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	return mux

}
