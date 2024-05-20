build:
	go build -o bin/e-commerce-api -v cmd/main.go

run: build
	bin/e-commerce-api

test: 
	go test -v ./...