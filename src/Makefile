# Variables
MOD_NAME=github.com/aldrickdev/go-htmx
binary=$(PWD)/main

# Basic Commands
build:
	go build -C cmd/app -o $(binary) -buildvcs=false 

run-build: build
	./app

run:
	go run cmd/app/app.go

init:
	go mod init $(MOD_NAME)
	go mod tidy
