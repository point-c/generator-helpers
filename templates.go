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

type Template interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

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

func Generate[T Template, D any](tmpl T, dot D, name, out string) {
	b := ExecTemplate[T, D](tmpl, dot, name)
	Check(os.WriteFile(out, b, os.ModePerm))
}

func ExecTemplate[T Template, D any](tmpl T, dot D, name string) []byte {
	b := Must(Exec(tmpl, dot, name))
	bf, err := GoFmt(b)
	if err != nil {
		slog.Error("failed to format code", "error", err)
		return b
	}
	return bf
}

func Exec[T Template, D any](tmpl T, d D, startTemplate string) ([]byte, error) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, startTemplate, d); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
