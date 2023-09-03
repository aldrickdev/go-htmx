# Variables
MOD_NAME=github.com/aldrickdev/go-htmx

# Basic Commands
build:
	go build -C cmd/app -buildvcs=false
	mv cmd/app/app ./

run-build: build
	./app

run:
	go run cmd/app/app.go

init:
	go mod init $(MOD_NAME)
	go mod tidy
