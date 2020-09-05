.PHONY: build
build:
	go build -v ./cmd/doublylinkedlist

.PHONY: run
run:
	go run ./cmd/doublylinkedlist

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build