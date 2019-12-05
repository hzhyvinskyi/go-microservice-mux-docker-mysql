.PHONY: run
run:
	go build -o ./bin/app ./cmd/webp-converter && ./bin/app

.PHONY: build
build:
	go build -o ./bin/app ./cmd/webp-converter

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: test-bench
test-bench:
	go test -bench=. ./...

.PHONY: test-cover
test-cover:
	go test -cover ./...

.DEFAULT_GOAL := run