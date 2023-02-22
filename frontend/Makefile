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
	@yarn run eslint --ext .js,.vue src

format:  ## Format code
	@yarn run eslint --ext .js,.vue src --fix

## @ Application
.PHONY: run docker
run:  ## Run project
	@yarn run serve

docker:  ## Build a docker image
	@docker build -t sorveteria-tres-estrelas-frontend .

