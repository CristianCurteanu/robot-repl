# include .env

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOGET=go get -u
GOBUILD=go build main.go
BUILDARCH=$(shell arch)

build:
	$(GOBUILD)

run:
	$(GOBUILD)
	./main

run-osx:
	GOOS=darwin GOARCH=amd64 $(GOBUILD)
	./main

run-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(BUILDARCH) $(GOBUILD)
	./main
