# Variables
MOD_NAME=github.com/aldrickdev/gin_api_template

# Basic Commands
build:
	go build -C cmd/app
	mv cmd/app/app ./

run: build
	./app

init:
	go mod init $(MOD_NAME)
	go mod tidy
