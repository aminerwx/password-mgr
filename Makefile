BINARY_NAME=password-mgr

build:
	@mkdir bin
	@GOARCH=amd64 GOOS=linux go build -o bin/$(BINARY_NAME)-linux main.go

run:
	@go run main.go

vault:
	@go run cmd/vault/main.go

clean:
	@go clean
	@rm -rf ./bin

dep:
	@go mod download

test:
	@go test ./...

coverage:
	@go test ./... -cover
