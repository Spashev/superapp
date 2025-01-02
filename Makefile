.DEFAULT_GOAL := help

help: ## Help message
	@echo "Please choose a task:"
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'


PROJECT_DIR=$(shell dirname $(realpath $(MAKEFILE_LIST))) ## Project directory

ifeq (manage,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

run: ## Run go server
	go run ./cmd/server/main.go
build: ## Build project
	go build -o ./bin/bookit ./cmd/server/main.go
test: ## Run all tests
	go test -v -race -timeout 30s ./...

# Flags
.PHONY: *
