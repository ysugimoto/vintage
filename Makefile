.PHONY: test

dev: generate test lint

test:
	go test ./...

lint:
	golangci-lint run

generate:
	cd cmd/generator && go run .
