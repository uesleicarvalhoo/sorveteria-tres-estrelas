COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

# Load environment variables from .env file
-include .env
export

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Tools
.PHONY: install-tools
install-tools:  ## Instal mockery, gofumpt, swago and golangci-lint
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/vektra/mockery/v22.20.0
	@go install github.com/swaggo/swag/cmd/swag@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0

## @ Linter
.PHONY: lint format
lint:  ## Run golangci-lint
	@golangci-lint run -v

format:  ## Format code
	@gofumpt -e -l -w .

## @ Application
.PHONY: swagger run
cmd/api/fiber/docs/*: $(wildcard cmd/api/main.go) $(wildcard cmd/api/fiber/handler/*.go) $(wildcard */entity.go) $(wildcard */dto.go) ## Generate swagger docs
	@swag init --generalInfo ./cmd/api/main.go --output ./cmd/api/fiber/docs

swagger: cmd/api/fiber/docs/*  ## Generate swagger docs

run: swagger  ## Run backend server
	@go run cmd/api/*.go

## @ Tests
.PHONY: test test/unit test/integration coverage clean-mocks generate-mocks
generate-mocks: clean-mocks  ## Generate mock files
	@mockery --dir product --output product/mocks --all
	@mockery --dir sales --output sales/mocks --all
	@mockery --dir user --output user/mocks --all
	@mockery --dir auth --output auth/mocks --all
	@mockery --dir cache --output cache/mocks --all
	@mockery --dir healthcheck --output healthcheck/mocks --all
	@mockery --dir payment --output payment/mocks --all
	@mockery --dir cashflow --output cashflow/mocks --all

clean-mocks:  ## Clean mock files
	@rm -rf product/mocks/*
	@rm -rf sales/mocks/*
	@rm -rf user/mocks/*
	@rm -rf auth/mocks/*
	@rm -rf cache/mocks/*
	@rm -rf healthcheck/mocks/*
	@rm -rf payment/mocks/*
	@rm -rf cashflow/mocks/*

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
