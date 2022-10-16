COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

# Load environment variables from .env file
-include .env
export

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_/-]+:.*?##/ { printf "\033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Linter
.PHONY: lint format
lint:  ## Run golangci-lint
	@golangci-lint run -v

format:  ## Format code
	@gofumpt -e -l -w .

## @ Application
.PHONY: swagger run
api/docs/*: $(wildcard api/main.go) $(wildcard api/handler/*.go)
	@swag init --generalInfo api/main.go --output ./api/docs

swagger: api/docs/*  ## Generate Swagger content

run: swagger  ## Run app
	@go run api/*.go

## @ Tests
.PHONY: test test/unit test/integration coverage clean-mocks generate-mocks
generate-mocks: clean-mocks  ## Generate mock files
	@mockery --dir usecase/popsicle --output usecase/popsicle/mocks --all
	@mockery --dir usecase/sales --output usecase/sales/mocks --all
	@mockery --dir usecase/user --output usecase/user/mocks --all
	@mockery --dir auth --output auth/mocks --all
	@mockery --dir cache --output cache/mocks --all

clean-mocks:  ## Clean mock files
	@rm -rf usecase/popsicle/mocks/*
	@rm -rf usecase/sales/mocks/*
	@rm -rf usecase/user/mocks/*
	@rm -rf auth/mocks/*
	@rm -rf cache/mocks/*

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
