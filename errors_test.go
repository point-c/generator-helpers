package generator_helpers

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/require"
	"log/slog"
	"testing"
)

func preserveExit(t *testing.T) func() {
	t.Helper()
	original := exit
	exit = func(code int) { t.Error(code) }
	return func() { t.Helper(); exit = original }
}

func TestCheck(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		defer preserveExit(t)()
		defer slog.SetDefault(slog.Default())
		var buf bytes.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
		var run bool
		exit = func(code int) { run = true; require.Equal(t, code, 1) }
		Check(errors.New("test"))
		require.True(t, run)
		require.NotEmpty(t, buf.Bytes())
	})

	t.Run("no error", func(t *testing.T) {
		defer preserveExit(t)()
		defer slog.SetDefault(slog.Default())
		var buf bytes.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
		exit = func(int) { t.Fail() }
		Check(nil)
		require.Empty(t, buf.Bytes())
	})
}

func TestMust(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		defer preserveExit(t)()
		defer slog.SetDefault(slog.Default())
		var buf bytes.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
		var run bool
		exit = func(code int) { run = true; require.Equal(t, code, 1) }
		require.Equal(t, 1, Must(1, errors.New("test")))
		require.True(t, run)
		require.NotEmpty(t, buf.Bytes())
	})

	t.Run("no error", func(t *testing.T) {
		defer preserveExit(t)()
		defer slog.SetDefault(slog.Default())
		var buf bytes.Buffer
		slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
		exit = func(int) { t.Fail() }
		require.Equal(t, 1, Must(1, nil))
		require.Empty(t, buf.Bytes())
	})
}
