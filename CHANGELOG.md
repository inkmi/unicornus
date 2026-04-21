# Changelog

All notable changes to this project are documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Godoc comments on the public API (`FormLayout`, `Choice`, `DataField`, `RenderContext`, `RenderForm`, `RenderView`, and related).
- `DIVRaw` render helper for trusted, already-escaped HTML fragments.
- CI workflow now runs `go vet`, `staticcheck`, `gocyclo`, `govulncheck`, `gosec`, and `gitleaks` in addition to build and test.

### Changed
- Upgraded `golang.org/x/net` to v0.53.0 and other dependencies to their latest minor versions.
- GitHub Actions workflow upgraded to `actions/checkout@v4` and `actions/setup-go@v5`; Go version is now read from `go.mod`.
- README: removed the "ALPHA - play, don't use" banner and the stale coverage badge.

### Removed
- Dead `TailwindOldTheme` implementation.
- Stray working files (`i.md`, `notailwind.html`, `tmp/`) that had crept into the repo.

### Fixed
- **Security**: HTML escaping is now applied to all user-controlled values flowing into attribute values and text nodes in the default render path (choice labels/values, field names, classes, styles, ids, placeholders, optgroup labels, validation error messages, and the anchor derived from group labels). Previously these were written raw, allowing XSS via struct-tag metadata or Choice data. If you relied on embedding pre-rendered HTML in labels or error messages, you must now pass it via `DIVRaw` or the existing `WriteUnsafeValue`/`NoEscape` path.
- Nil dereference in `RenderElementWithErrors` when the named element does not exist.
- Defensive guards around `strings.SplitN` result in `findByName`.
