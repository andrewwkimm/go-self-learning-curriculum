help:
	cat Makefile

################################################################################

build:
	go mod download
	make reformat
	make lint
	make type_check
	go build ./...
	make test

lint:
	go vet ./...
	golangci-lint run --fix ./...

reformat:
	go fmt ./...

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install honnef.co/go/tools/cmd/staticcheck

test:
	go test -v -race ./...

type-check:
	staticcheck ./...

################################################################################

clean:
	go clean -cache
	go clean -fuzzcache

################################################################################

.PHONY: \
	build \
	help \
	lint \
	reformat \
	test \
	type-check