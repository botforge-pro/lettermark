STATICCHECK_VERSION := v0.8.1
STATICCHECK := $(shell go env GOPATH)/bin/staticcheck

.PHONY: install-tools format lint comments test-build test build sync-corpus

sync-corpus:
	cp cases.yaml ../lettermark-swift/Tests/LettermarkTests/Resources/cases.yaml
	cp cases.yaml ../lettermark-kotlin/src/test/resources/cases.yaml

install-tools:
	go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)
	python3 -m pip install --quiet --upgrade git+https://github.com/botforge-pro/commentcensor.git

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
