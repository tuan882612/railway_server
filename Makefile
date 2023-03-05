run:
	GOARCH=amd64 GOOS=linux ./bin/main-linux
	GOARCH=amd64 GOOS=darwin ./bin/main-mac
	GOARCH=amd64 GOOS=windows ./bin/main-win

dev:
	go run ./cmd/api/main.go

build:
	GOARCH=amd64 GOOS=linux go build -o ./bin/main-linux ./cmd/api/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./bin/main-mac ./cmd/api/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/main-win ./cmd/api/main.go

testAll:
	go clean -testcache
	go test -v ./test/...

testConfig:
	go clean -testcache
	go test -v ./test/config

testHandlers:
	go clean -testcache
	go test -v ./test/handlers
