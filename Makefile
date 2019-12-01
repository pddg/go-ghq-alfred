SHELL := /bin/bash
BIN_NAME := dist/go-ghq-alfred
WF_NAME := ghq-alfred.alfredworkflow
ASSETS := $(shell find -f ./resources)
TESTDIR := testdir

$(TESTDIR):
	mkdir -p $(TESTDIR)/{data,cache}

test: $(TESTDIR)
	env alfred_workflow_bundleid=testid \
		alfred_workflow_cache=$(TESTDIR)/cache \
		alfred_workflow_data=$(TESTDIR)/data \
		go test ./... -v

$(BIN_NAME): main.go
	go build -o $(BIN_NAME) .

build: $(BIN_NAME)

$(WF_NAME): $(BIN_NAME) $(ASSETS)
	if [ ! -d dist ]; then \
		mkdir dist/; \
	fi
	cp -r resources/* dist/
	cd dist && zip -r ../$(WF_NAME) ./*

dist: $(WF_NAME)
