SHELL:=/bin/bash

.PHONY: run lint build build-linux

# go get github.com/pilu/fresh 全局安装fresh命令,热更新代码
run:
	source ./env.sh && fresh -c dev-run.conf

lint:
	go fmt ./internal/... && go fmt ./main.go

build:
	go build -o cluture

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cluture-linux
