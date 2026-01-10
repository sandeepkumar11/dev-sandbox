package config

import "os"

type Config struct {
	Port    string
	BaseURL string
}

func Load() (*Config, error) {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:" + port
	}

	return &Config{
		Port:    port,
		BaseURL: baseURL,
	}, nil
}
