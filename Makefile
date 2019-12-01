SHELL := /bin/bash
BIN_NAME := dist/go-ghq-alfred
WF_NAME := ghq-alfred.alfredworkflow

test:
	go test ./... -v

$(BIN_NAME): main.go
	go build -o $(BIN_NAME) .

build: $(BIN_NAME)

$(WF_NAME): build
	if [ ! -d dist ]; then \
		mkdir dist/; \
	fi
	cp -r resources/* dist/
	cd dist && zip -r ../$(WF_NAME) ./*

dist: $(WF_NAME)
