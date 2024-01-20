package generator_helpers

import (
	"bytes"
	html_template "html/template"
	"io"
	"io/fs"
	"log/slog"
	"os"
	text_template "text/template"
)

// Template interface abstracts the common functionality for text and HTML templates.
type Template interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

// NewTemplate initializes a new template, either text or HTML, with provided functions and file system.
func NewTemplate[T Template](templates fs.FS, funcs map[string]any) T {
	var t Template
	switch any(*new(T)).(type) {
	case *text_template.Template:
		t = Must(text_template.New("").Funcs(funcs).ParseFS(templates, "*"))
	case *html_template.Template:
		t = Must(html_template.New("").Funcs(funcs).ParseFS(templates, "*"))
	}
	return t.(T)
}

// Generate creates and writes a file using a template and data.
func Generate[T Template, D any](tmpl T, dot D, name, out string) {
	b := ExecTemplate[T, D](tmpl, dot, name)
	Check(os.WriteFile(out, b, os.ModePerm))
}

// ExecTemplate executes a template and returns the generated content as a byte slice.
// It assumes the output is go source code and formats the output.
func ExecTemplate[T Template, D any](tmpl T, dot D, name string) []byte {
	b := Must(Exec(tmpl, dot, name))
	bf, err := GoFmt(b)
	if err != nil {
		slog.Error("failed to format code", "error", err)
		return b
	}
	return bf
}

// Exec executes a template with the provided data and returns the result as a byte slice.
func Exec[T Template, D any](tmpl T, d D, startTemplate string) ([]byte, error) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, startTemplate, d); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
