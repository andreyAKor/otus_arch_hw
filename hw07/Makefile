GOBASE=$(shell pwd)
GOBIN=$(GOBASE)

build:
	go build -ldflags="-s -w" -o '$(GOBIN)/hw07' ./main.go || exit

run:
	go build -o '$(GOBIN)/hw07' ./main.go
	'$(GOBIN)/hw07' -limit 20 -algo fibonacci -out out_fibonacci.txt
	'$(GOBIN)/hw07' -limit 20 -algo binet -out out_binet.txt

test:
	go test -v -count=1 -race -timeout=60s ./...

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

install:
	go mod download

generate:
	go generate ./...

.PHONY: build
