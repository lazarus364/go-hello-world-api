# Go Hello World API

A minimal HTTP server in Go that returns JSON at `/`.

## Requirements
- Go 1.21+

## Run
```bash
go run .
```
Visit: http://localhost:8080/

## Test with cURL
```bash
curl http://localhost:8080/
```

Expected JSON with a message and timestamp.

## Build binary
```bash
go build -o app .
./app        # Linux/macOS
app.exe      # Windows
```

## Troubleshooting
- Ensure `go version` works.
- Run `go mod init example.com/go-hello-world-api` if missing.
- Change port in `main.go` if 8080 is busy.
