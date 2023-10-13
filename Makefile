.PHONY: test

BUILD_VERSION=$(or ${VERSION}, dev)

dev: generate test lint

test:
	go test ./...

lint:
	golangci-lint run

generate:
	cd cmd/generator && go run .

linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
			 -ldflags "-X main.version=$(BUILD_VERSION)" \
			 -o dist/vintage-linux-amd64 ./cmd/vintage
	cd ./dist/ && cp ./vintage-linux-amd64 ./vintage && tar cfz vintage-linux-amd64.tar.gz ./vintage

darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build \
			 -ldflags "-X main.version=$(BUILD_VERSION)" \
			 -o dist/vintage-darwin-amd64 ./cmd/vintage
	cd ./dist/ && cp ./vintage-darwin-amd64 ./vintage && tar cfz vintage-darwin-amd64.tar.gz ./vintage

darwin_arm64:
	GOOS=darwin GOARCH=arm64 go build \
			 -ldflags "-X main.version=$(BUILD_VERSION)" \
			 -o dist/vintage-darwin-arm64 ./cmd/vintage
	cd ./dist/ && cp ./vintage-darwin-arm64 ./vintage && tar cfz vintage-darwin-arm64.tar.gz ./vintage

all: linux darwin_amd64 darwin_arm64

clean:
	rm ./dist/vintage-*
