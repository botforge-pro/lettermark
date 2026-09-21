STATICCHECK_VERSION := v0.8.1
STATICCHECK := $(shell go env GOPATH)/bin/staticcheck

.PHONY: tools format lint test-build test build

tools:
	go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

format:
	gofmt -w .

lint:
	go vet ./...
	$(STATICCHECK) ./...
	gofmt -l . | (! grep .)

test-build:
	go build ./...
	go test -run '^$$' ./...

test:
	go test -count=1 ./...

build:
	go build ./...
