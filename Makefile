PREFIX?=$(shell pwd)

.PHONY: clean all fmt vet lint build test
.DEFAULT: all

all: clean fmt vet lint build test
dist: clean build

APP_NAME := main

GOLINT := $(shell which golint || echo '')

PKGS := $(shell go list ./... | grep -v ^github.com/everywan/go-web/vendor/)

clean:
	@echo "+ $@"
	@rm -rf "${PREFIX}/bin/"

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)

vet:
	@echo "+ $@"
	@go vet $(PKGS)

lint:
	@echo "+ $@"
	$(if $(GOLINT), , \
		$(error Please install golint: `go get -u github.com/golang/lint/golint 需翻墙！`))
	@test -z "$$($(GOLINT) ./... 2>&1 | grep -v ^vendor/ | tee /dev/stderr)"

test:
	@echo "+ $@"
	@go test -test.short $(PKGS)

build:
	@echo "+ $@"
	@mkdir "bin"
	@go build -v -o ${APP_NAME}
	@mv ${APP_NAME} "${PREFIX}/bin/"
