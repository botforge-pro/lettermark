STATICCHECK_VERSION := v0.8.1
COMMENTCENSOR_VERSION := v0.1.0
STATICCHECK := $(shell go env GOPATH)/bin/staticcheck

.PHONY: tools format lint comments test-build test build

tools:
	go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)
	python3 -m pip install --quiet \
		git+https://github.com/botforge-pro/commentcensor.git@$(COMMENTCENSOR_VERSION)

format:
	gofmt -w .

lint: comments
	go vet ./...
	$(STATICCHECK) ./...
	gofmt -l . | (! grep .)
	go mod tidy -diff

comments:
	commentcensor .

test-build:
	go build ./...
	go test -run '^$$' ./...

test:
	go test -count=1 ./...

build:
	go build ./...
