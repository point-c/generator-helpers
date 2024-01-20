# generator_helpers

[![Go Reference](https://img.shields.io/badge/godoc-reference-%23007d9c.svg)](https://point-c.github.io/generator-helpers)

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

### Go Code Formatting

The `generator_helpers` package offers functionalities to format Go source code. Below is an example of how to use these functions:

#### Formatting Go Source Code

To format a Go source code byte slice:

```go
sourceCode := []byte("package main\nimport \"fmt\"\nfunc main() {fmt.Println(\"hello world\")}")
formattedCode, err := generator_helpers.GoFmt(sourceCode)
if err != nil {
    // Handle formatting error
}
// Use formattedCode...
```

#### Formatting Source Code from an `io.Reader`

If you have Go source code in an `io.Reader`, you can format it as follows:

```go
var reader io.Reader = ... // your io.Reader with Go source code
formattedCode, err := generator_helpers.GoFmtReader(reader)
if err != nil {
    // Handle formatting error
}
// Use formattedCode...
```

## Testing

The package includes tests that demonstrate its functionality. Use Go's testing tools to run the tests:

```bash
go test
```

## Godocs

To regenerate godocs:

```bash
go generate -tags docs ./...
```
