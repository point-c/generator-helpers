package generator_helpers

import (
	"bytes"
	"embed"
	"github.com/stretchr/testify/require"
	html_template "html/template"
	"log/slog"
	"os"
	"path/filepath"
	"testing"
	text_template "text/template"
)

//go:embed *.gotmpl
var testFS embed.FS

func TestNewTemplate(t *testing.T) {
	defer preserveExit(t)()
	defer slog.SetDefault(slog.Default())
	var lbuf bytes.Buffer
	slog.SetDefault(slog.New(slog.NewTextHandler(&lbuf, nil)))
	funcs := map[string]any{
		"testFunc": func() string { return "test" },
	}

	t.Run("text", func(t *testing.T) {
		exit = func(i int) { t.Error(i) }
		textTmpl := NewTemplate[*text_template.Template](testFS, funcs)
		require.Empty(t, lbuf.Bytes(), lbuf.String())
		require.NotNil(t, textTmpl)
	})

	t.Run("html", func(t *testing.T) {
		exit = func(i int) { t.Error(i) }
		htmlTmpl := NewTemplate[*html_template.Template](testFS, funcs)
		require.Empty(t, lbuf.Bytes(), lbuf.String())
		require.NotNil(t, htmlTmpl)
	})
}

func TestGenerate(t *testing.T) {
	defer preserveExit(t)()
	defer slog.SetDefault(slog.Default())
	var lbuf bytes.Buffer
	slog.SetDefault(slog.New(slog.NewTextHandler(&lbuf, nil)))
	tmpl := NewTemplate[*text_template.Template](testFS, nil)
	data := struct{ Name string }{"Test"}
	outFile := filepath.Join(t.TempDir(), "output.txt")
	// Clean up
	defer func() { require.NoError(t, os.Remove(outFile)) }()

	Generate(tmpl, data, "test.gotmpl", outFile)
	require.Empty(t, lbuf.Bytes(), lbuf.String())

	// Check if file is generated
	_, err := os.Stat(outFile)
	require.NoError(t, err)

}

func TestExecTemplate(t *testing.T) {
	defer preserveExit(t)()
	defer slog.SetDefault(slog.Default())
	var lbuf bytes.Buffer
	slog.SetDefault(slog.New(slog.NewTextHandler(&lbuf, nil)))
	tmpl := NewTemplate[*text_template.Template](testFS, nil)

	data := struct{ Name string }{"Test"}
	t.Run("no error", func(t *testing.T) {
		result := ExecTemplate(tmpl, data, "test.gotmpl")
		require.NotEmpty(t, result)
		require.Empty(t, lbuf.Bytes())
	})

	t.Run("error", func(t *testing.T) {
		data.Name = "{}"
		result := ExecTemplate(tmpl, data, "test.gotmpl")
		require.NotEmpty(t, result)
		require.NotEmpty(t, lbuf.Bytes())
	})
}

func TestExec(t *testing.T) {
	defer preserveExit(t)()
	defer slog.SetDefault(slog.Default())
	var lbuf bytes.Buffer
	slog.SetDefault(slog.New(slog.NewTextHandler(&lbuf, nil)))

	exit = func(int) { t.Fatal() }
	tmpl := NewTemplate[*text_template.Template](testFS, nil)

	t.Run("no error", func(t *testing.T) {
		lbuf.Reset()
		data := struct{ Name string }{"Test"}
		buf, err := Exec(tmpl, data, "test.gotmpl")
		require.NoError(t, err)
		require.NotEmpty(t, buf)
		require.Empty(t, lbuf.Bytes())
	})

	t.Run("error", func(t *testing.T) {
		lbuf.Reset()
		buf, err := Exec(tmpl, struct{}{}, "test.gotmpl")
		require.Error(t, err)
		require.Empty(t, buf)
		require.Empty(t, lbuf.Bytes())
	})
}
