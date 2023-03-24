build:
	go build -o app cmd/main.go

start:
	./app

lint:
	golangci-lint .golangci.yml