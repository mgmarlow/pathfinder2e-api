BINARY_NAME=upserter

build:
	go build -o bin/$(BINARY_NAME).exe cmd/upserter/main.go

test:
	go test -v ./...

run:
	go run cmd/upserter/main.go
