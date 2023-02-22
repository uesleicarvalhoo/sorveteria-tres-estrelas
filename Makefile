COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

# Load environment variables from .env file
-include .env
export

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Tools
.PHONY: install-tools
install-tools:  ## Instal mockery, gofumpt, swago and golangci-lint
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0

## @ Linter
.PHONY: lint format
lint:  ## Run golangci-lint
	@golangci-lint run -v

format:  ## Format code
	@gofumpt -e -l -w .

## @ Application
.PHONY: swagger run compose
cmd/api/fiber/http/docs/*: $(wildcard cmd/api/http/fiber/main.go) $(wildcard cmd/api/http/fiber/handler/*.go) $(wildcard */entity.go) $(wildcard */repository.go) $(wildcard */dto.go) ## Generate swagger docs
	@swag init --generalInfo ./cmd/api/main.go --output ./cmd/api/fiber/docs

swagger: cmd/api/fiber/http/docs/*  ## Generate swagger docs

run: swagger  ## Run app
	@go run cmd/api/*.go

compose:  ## Init containers with dev dependencies
	@docker compose build && docker compose up -d

## @ Tests
.PHONY: test test/unit test/integration coverage clean-mocks generate-mocks
generate-mocks: clean-mocks  ## Generate mock files
	@mockery --dir products --output products/mocks --all
	@mockery --dir sales --output sales/mocks --all
	@mockery --dir users --output users/mocks --all
	@mockery --dir auth --output auth/mocks --all
	@mockery --dir cache --output cache/mocks --all
	@mockery --dir healthcheck --output healthcheck/mocks --all

clean-mocks:  ## Clean mock files
	@rm -rf products/mocks/*
	@rm -rf sales/mocks/*
	@rm -rf users/mocks/*
	@rm -rf auth/mocks/*
	@rm -rf cache/mocks/*
	@rm -rf healthcheck/mocks/*

test:  ## Run tests all tests
	@go test ./... -race -v -count=1 -tags="all" -coverprofile=$(COVERAGE_OUTPUT)

test/unit:  ## Run unit tests
	@go test ./... -race -v -count=1 -tags="unit" -coverprofile=$(COVERAGE_OUTPUT)

test/integration:  ## Run integration tests
	go test ./... -race -v -count=1 -tags="integration" -coverprofile=$(COVERAGE_OUTPUT)

coverage: test ## Run tests, make coverage report and display it into browser
	@go test ./... -race -v -cover
	@go tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)
	@wslview ./$(COVERAGE_HTML) || xdg-open ./$(COVERAGE_HTML) || powershell.exe Invoke-Expression ./$(COVERAGE_HTML)

## @ Clean
.PHONY: clean clean_coverage_cache
clean_coverage_cache:
	@rm -rf $(COVERAGE_OUTPUT)
	@rm -rf $(COVERAGE_HTML)

clean: clean_coverage_cache ## Remove Cache files
