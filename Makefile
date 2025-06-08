APP_NAME = url-shortener
PKG_LIST := $(shell go list ./...)

.DEFAULT_GOAL := help

## Show help
.PHONY: help
help:
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## ' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'
	@echo ""

test: 
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

cover-html: test 
	go tool cover -html=coverage.out -o coverage.html
	@echo "View coverage at: file://$(CURDIR)/coverage.html"

tidy: 
	go mod tidy
	go fmt ./...
	go vet ./...

clean: 
	rm -f coverage.out coverage.html

run: 
	go run main.go

build: 
	go build -o $(APP_NAME) 

docker-build: 
	docker compose build

docker-up: 
	docker compose up


docker-down: 
	docker compose down
