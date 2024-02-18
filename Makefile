# https://maex.me/2018/02/dont-fear-the-makefile/

test:
	gotestsum ./...

staticcheck:
	staticcheck ./pkg

lint: staticcheck
	golangci-lint run
	go vet ./...

# https://github.com/fdaines/arch-go
arch:
	arch-go

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

example:
	go build -o bin ./...
	./bin/example

air:
	air
