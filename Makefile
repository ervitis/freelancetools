.PHONY: tests

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

tests: ## Run tests with race detection
	command -v testtime &> /dev/null || echo "I require testtime but it's not installed.  Installing."; go install github.com/tenntenn/testtime/cmd/testtime@latest
	go test -race -v -overlay=`testtime` ./...

run: ## Run the application
	go run ./invoices/cmd/main.go

install: ## Install dependencies
	go mod download