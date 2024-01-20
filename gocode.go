package generator_helpers

import (
	"go/format"
	"io"
)

// GoFmt formats Go source code, returning the formatted code or the original and an error.
func GoFmt(src []byte) (dst []byte, err error) {
	dst, err = format.Source(src)
	if err != nil {
		dst = src
	}
	return
}

// GoFmtReader functions similarly to [GoFmt] but accepts an io.Reader as input.
// It reads the source code from the reader, formats it, and returns the formatted code and any error encountered.
func GoFmtReader(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return GoFmt(b)
}
