.PHONY: all

all: test

test:
	@go test -failfast -v ./internal/...
