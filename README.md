# PdfToImages

A PDF to PNG converter written in Go.

## Requirements

- Go >= 1.25.0
- Make

## Quick Start

### Build
```bash
make build
````

```bash
make clean
```

### Run Example

```bash
make run
```

### Run Tests

```bash
make test
```

### Install Binary to GOPATH/bin

```bash
make install
```

### Manage Dependencies

```bash
make deps
```

### Format and Vet Code

```bash
make fmt
make vet
```

## Continuous Integration (CI)

GitHub Actions automatically build and run all tests on every push and pull request to the `main` branch.

## Makefile Targets

* `build` — build the binary
* `clean` — clean the project
* `run` — build and run the example
* `test` — run all tests (unit + E2E)
* `install` — install the binary to GOPATH/bin
* `deps` — update dependencies
* `fmt` — format code
* `vet` — check code for potential issues
