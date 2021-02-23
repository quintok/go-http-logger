.DEFAULT_GOAL := build

BIN_FILE=server

build: server.go
	@go build -o "${BIN_FILE}"

run:
	./"${BIN_FILE}"