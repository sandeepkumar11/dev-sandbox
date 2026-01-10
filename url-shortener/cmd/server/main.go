package main

import (
	"log"
	"net/http"

	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/app"
	"github.com/sandeepkumar11/dev-sandbox/url-shortner/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	router := app.SetupRouter(cfg)

	log.Printf("Starting server on :%s\n", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
