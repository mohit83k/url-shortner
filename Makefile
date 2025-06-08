# Project settings
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

## Run tests with coverage
test: ## Run unit tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

## Generate HTML coverage report
cover-html: test ## Open HTML coverage report
	go tool cover -html=coverage.out -o coverage.html
	@echo "View coverage at: file://$(CURDIR)/coverage.html"

## Format, vet, tidy
tidy: ## Tidy and format code
	go mod tidy
	go fmt ./...
	go vet ./...

## Clean coverage files
clean: ## Remove generated files
	rm -f coverage.out coverage.html

## Run app locally (host)
run: ## Run app directly using go run
	go run main.go

## Build binary
build: ## Build local binary
	go build -o $(APP_NAME) main.go

## Docker: Build image
docker-build: ## Build docker image using docker-compose
	docker compose build

## Docker: Run container
docker-up: ## Run docker container using docker-compose
	docker compose up

## Docker: Stop container
docker-down: ## Stop and remove docker container
	docker compose down
