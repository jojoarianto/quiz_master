.PHONY: migrate-schema run

migrate-schema: ## migrate database
	go run interface/cli/migration/main.go 

test: ## run go test
	go test ./... 

coverage: ## Generate global code coverage report
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

preview-coverage: ## Preview test coverage
	go tool cover -html=coverage.txt
	
run: ## run on development mode
	go run main.go

build: ## build binary file
	go build -o quiz_master main.go

help: ## display help command
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
