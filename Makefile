.PHONY: all

all: test

test:
	@clear && go test -failfast -v -cover ./internal/...
