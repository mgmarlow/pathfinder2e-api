build:
	go build -o bin/main.exe main.go parser.go monster.go

run:
	go run main.go parser.go monster.go
