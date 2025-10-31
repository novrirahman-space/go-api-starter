.PHONY: run test lint build docker

run:
	@APP_ENV=development HTTP_ADDR=:8080 go run ./cmd/server

test:
	@go test ./... -race -coverprofile=coverage.out

lint:
	@golangci-lint run

build:
	@CGO_ENABLED=0 go build -o bin/server ./cmd/server

docker:
	@docker build -t ghcr.io/novrirahman-space/go-api-starter:latest .
