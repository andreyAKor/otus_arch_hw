GOBASE=$(shell pwd)
GOBIN=$(GOBASE)

build:
	go build -ldflags="-s -w" -o '$(GOBIN)/hw04' ./main.go || exit

run:
	go build -o '$(GOBIN)/hw04' ./main.go
	'$(GOBIN)/hw04' -algo 1 -in in.txt -out out_selection.txt
	'$(GOBIN)/hw04' -algo 2 -in in.txt -out out_insertion.txt
	'$(GOBIN)/hw04' -algo 3 -in in.txt -out out_merge.txt

test:
	go test -v -count=100 -race -timeout=60s ./...

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

install:
	go mod download

generate:
	go generate ./...

.PHONY: build
