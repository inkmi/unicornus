# https://maex.me/2018/02/dont-fear-the-makefile/

test:
	gotestsum ./...

staticcheck:
	staticcheck ./pkg

lint: staticcheck
	golangci-lint run
	go vet ./...

audit:
	go list -json -deps ./... | nancy sleuth --loud

sec: audit
	gosec  .
	govulncheck ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

upgrade-deps:
	go get -u ./...
	go mod tidy
	gotestsum ./...

all:   lint sec  test

examples:
	go build -o bin ./...
	./bin/example

air:
	air

doc:
	@../eris/bin/eris --dir docs/ --out doc.md
	@cat README_header.md doc.md > README.md