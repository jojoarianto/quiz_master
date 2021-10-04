.PHONY: migrate-schema run

migrate-schema:
	go run interface/cli/migration/main.go 

test: 
	go test ./... 

run:
	go run main.go
