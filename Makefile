COVERAGE_OUTPUT=coverage.output
COVERAGE_HTML=coverage.html

# Load environment variables from .env file
-include .env
export

## @ Help
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "\033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

## @ Linter
.PHONY: lint format
lint:
	@golangci-lint run -v

format:
	@gofumpt -e -l -w .

## @ Tests
.PHONY: test coverage
generate-mocks:  # Generate mock files
	@mockery --dir popsicle --output popsicle/mocks --all
	@mockery  --dir sales --output sales/mocks --all

clean-mocks:  # Clean mock files
	@rm popsicle/mocks/*
	@rm sales/mocks/*

test: clean-mocks generate-mocks ## Run tests of project
	@go test ./... -race -v -count=1 -coverprofile=$(COVERAGE_OUTPUT)

coverage: test ## Run tests, make report and open into browser
	@go test ./... -race -v -cover
	@go tool cover -html=$(COVERAGE_OUTPUT) -o $(COVERAGE_HTML)
	@wslview ./$(COVERAGE_HTML) || xdg-open ./$(COVERAGE_HTML) || powershell.exe Invoke-Expression ./$(COVERAGE_HTML)

## @ Clean
.PHONY: clean clean_coverage_cache
clean_coverage_cache:
	@rm -rf $(COVERAGE_OUTPUT)
	@rm -rf $(COVERAGE_HTML)

clean: clean_coverage_cache ## Remove Cache files
