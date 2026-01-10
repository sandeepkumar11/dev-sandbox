# URL Shortener (Go)

A URL shortener built using Go with:
- Clean architecture
- Dependency Injection
- In-memory repository (swappable)
- net/http
- Thread-safe storage

## APIs

### POST /shorten
```json
{
  "url": "https://example.com"
}
```

## How to run the APP
```bash
cd url-shortener
export APP_PORT=8080
export BASE_URL=http://localhost:8080
go run cmd/server/main.go

curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://stackoverflow.com/questions"}'

curl -v http://localhost:8080/a1B2c3D4
```

