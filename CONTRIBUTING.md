# Contributing

Thanks for your interest in Unicornus.

## Development

```sh
make test       # run the test suite
make lint       # go vet + staticcheck + golangci-lint + nilaway + gocyclo
make sec        # gosec + govulncheck
make secrets    # gitleaks scan
make check      # all of the above
make coverage-check   # enforces the coverage gate
```

A Go toolchain matching `go.mod` is required.

## Pull requests

- Run `make check` locally before opening a PR.
- Keep changes focused. Large refactors and new features are easier to review as a series of small PRs.
- Add or update tests for any behavior change.
- When touching the render path, include an HTML-escaping test if the change affects attribute values or text nodes.

## Reporting security issues

Please do not open a public issue for suspected security bugs. Email the
maintainer listed in `README.md` instead.

## License

By contributing, you agree that your contributions will be licensed under
the MIT License (see `LICENSE`).
