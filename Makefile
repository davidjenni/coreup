SHELL="/bin/bash"

default: build validate test

OUT := coreUp
# PKGS := $(shell go list ./... | grep -vF /vendor/)
PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

build:
	go fmt ./...
	go build -o $(OUT) ./cmd/coreUp

validate: build lint vet
	golint $(PKGS)

lint:
	@for file in $(GO_FILES) ;  do \
		golint $$file ; \
	done

vet:
	go vet $(PKG_LIST)

test:
	go test -v ./...

install:
	go install ./cmd/coreUp

clean:
	-@rm $(OUT)
	-@rm coverage.txt
	-@rm debug.test

.PHONY: build validate lint vet test install clean