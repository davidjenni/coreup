SHELL="/bin/bash"

default: build validate test

OUT := coreUp
# PKGS := $(shell go list ./... | grep -vF /vendor/)
PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

get-tools:
	go get -u github.com/golang/lint/golint honnef.co/go/tools/cmd/megacheck

build:
	go fmt ./...
	go build -o $(OUT) ./cmd/coreUp

validate: build lint vet megacheck
	golint $(PKG_LIST)

lint:
	@for file in $(GO_FILES) ;  do \
		golint $$file ; \
	done

vet:
	go vet $(PKG_LIST)

# see: https://staticcheck.io
megacheck:
	megacheck $(PKG_LIST)

test:
	go test -v ./...

install:
	go install ./cmd/coreUp

clean:
	-@rm $(OUT)
	-@rm coverage.txt
	-@rm debug.test

.PHONY: build validate lint vet test install clean get-tools megacheck