#!/usr/bin/env bash

# aggregate test coverage results for codecov.io:
#   run tests with coverage over all non-vendor packages
#   merge all coverage files into single file
set -e
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
