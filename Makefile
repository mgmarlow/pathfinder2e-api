build:
	go build -o bin/upserter.exe cmd/upserter/main.go

run:
	go run cmd/upserter/main.go src/parser.go src/monster.go
