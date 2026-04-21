# https://maex.me/2018/02/dont-fear-the-makefile/
.PHONY: all build test coverage coverage-check lint sec secrets check clean upgrade-deps examples air doc tokens

VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

build:
	goimports -w .
	go build ./...

test:
	gotestsum ./...

coverage:
	gotestsum -- -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

coverage-check: coverage
	@go tool cover -func=coverage.out | awk '/^total:/{gsub(/%/,"",$$NF); printf "Total coverage: %s%%\n", $$NF; if ($$NF+0 < 45.0) {print "FAIL: below 45% threshold"; exit 1} else {print "OK: meets 45% threshold"}}'

lint:
	go vet ./...
	staticcheck ./...
	golangci-lint run ./...
	nilaway ./...
	gocyclo -over 15 .

sec:
	gosec ./...
	govulncheck ./...

secrets:
	gitleaks git -v

check: lint sec secrets

clean:
	go clean -cache -i
	rm -f coverage.out cover.out cover.html

upgrade-deps:
	go get -u ./...
	go mod tidy
	gotestsum ./...

tokens:
	@find . -name '*.go' ! -path './vendor/*' -exec cat {} + | wc -w | awk '{printf "%d words (~%d tokens)\n", $$1, int($$1 * 1.3)}'

all: lint sec test

examples:
	go build -o bin ./...
	./bin/example

air:
	air

doc:
	@../eris/bin/eris --dir docs/ --out doc.md --leanpub=false
	@cat README_header.md doc.md > README.md
