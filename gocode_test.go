package generator_helpers

import (
	"errors"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"testing/iotest"
)

const (
	testGoFmtInput    = "package main\nimport \"fmt\"\nfunc main() {\nfmt.Println(`Hello World`)}"
	testGoFmtExpected = "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(`Hello World`)\n}\n"
)

func TestGoFmt(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		result, err := GoFmt([]byte(testGoFmtInput))
		require.NoError(t, err)
		require.Equal(t, testGoFmtExpected, string(result))
	})
	t.Run("error", func(t *testing.T) {
		input := "package mainimport fmtfunc main() {fmt.Println(`Hello World`)}"
		result, err := GoFmt([]byte(input))
		require.Error(t, err)
		require.Equal(t, string(result), input)
	})
}

func TestGoFmtReader(t *testing.T) {
	t.Run("reader", func(t *testing.T) {
		result, err := GoFmtReader(strings.NewReader(testGoFmtInput))
		require.NoError(t, err)
		require.Equal(t, testGoFmtExpected, string(result))
	})
	t.Run("error", func(t *testing.T) {
		_, err := GoFmtReader(iotest.ErrReader(errors.New("")))
		require.Error(t, err)
	})
}
