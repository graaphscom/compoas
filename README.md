[![CI](https://github.com/graaphscom/compoas/actions/workflows/ci.yml/badge.svg)](https://github.com/graaphscom/compoas/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/graaphscom/compoas/branch/master/graph/badge.svg?token=T8UAQ93Z3T)](https://codecov.io/gh/graaphscom/compoas)
[![Go Reference](https://pkg.go.dev/badge/github.com/graaphscom/compoas.svg)](https://pkg.go.dev/github.com/graaphscom/compoas)

# compoas

Library for building, composing and serving
[OpenAPI Specification](https://github.com/OAI/OpenAPI-Specification)
(aka Swagger).

## Features

This lib provides:

- golang structs which reflect OpenAPI Specification (check Limitations),
- embedded [Swagger UI](https://github.com/swagger-api/swagger-ui) (through `go:embed`) and `http.Handler`
  for serving it,
- functions for merging OpenAPI Specifications and dumping them into a file

## Limitations

- included golang structs not cover entire OpenAPI Specification
- Swagger UI cannot be customized (only setting specification URL is possible)

## Installation

```
go get github.com/graaphscom/compoas
```

## Usage