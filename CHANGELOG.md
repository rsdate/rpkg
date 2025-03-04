# Changelog
## [0.0.2a1] - 2025-03-03

_Notice: This is an alpha release of v0.0.2_

### Added:
- Docker functionality for building the project into an image
- Github Workflows (no Docker image published yet)
### Changed:
- Moved excess functions in `./cmd/` to a file called `functions.go`
- Standardized error handling with the `RPKG_PANICMODE` environment variable and the `execStep()` function

## [0.0.1] - 2025-02-20

_First Release._