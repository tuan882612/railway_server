run:
	./bin/main

dev:
	go run ./cmd/api/main.go

build:
	go build -o ./bin ./cmd/api/main.go

testing:
	go clean -testcache
	go test -v ./test