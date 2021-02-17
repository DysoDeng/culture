SHELL:=/bin/bash

.PHONY: run lint build build-linux

run:
	source ./env.sh && go run ./cmd/main.go

lint:
	go fmt ./internal/... && go fmt ./cmd/main.go

build:
	cd cmd && go build -o cluture

build-linux:
	cd cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cluture-linux
