# generator_helpers

[![Go Reference](https://img.shields.io/badge/godoc-reference-%23007d9c.svg)](https://point-c.github.io/generator_helpers)

## Overview

`generator_helpers` is a Go package designed to provide utility functions for common operations such as template generation, YAML processing, and error handling in Go applications.

## Installation

To use `generator_helpers` in your Go project, you need to install it as a dependency:

```bash
go get github.com/point-c/generator_helpers
```

## Usage

Below are some examples of how to use the `generator_helpers` package:

### Unmarshalling YAML Data

To unmarshal YAML data from a file:

```go
var config MyConfig
err := generator_helpers.UnmarshalYAML("config.yaml", &config)
if err != nil {
    // Handle error
}
```

### Template Generation

Generating text or HTML from templates:

```go
tmpl, err := generator_helpers.NewTemplate[*text_template.Template](templateFS, funcs)
if err != nil {
    // Handle error
}

var data MyData
generator_helpers.Generate(tmpl, data, "templateName", "output.txt")
```

### Error Handling

Simplifying error handling with `Must` and `Check`:

```go
getValue := func() (int, error) { return 1, errors.New("error") }
value := generator_helpers.Must(getValue())
generator_helpers.Check(err)
```

## Godocs

To regenerate godocs:

```bash
go generate -tags docs ./...
```
