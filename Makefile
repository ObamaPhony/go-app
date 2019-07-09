SHELL := /bin/sh

GO := go
GO_OPTS :=

.POSIX:
.PHONY: all build clean test ci

all: deps build test

deps:
	go mod download
	go mod vendor

build:
	${GO} ${GO_OPTS} build ./cmd/...

clean:
	${GO} ${GO_OPTS} clean -x

test:
	${GO} ${GO_OPTS} test -v ./...

ci: clean deps build test
