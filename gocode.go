package generator_helpers

import (
	"go/format"
	"io"
)

func GoFmt(src []byte) (dst []byte, err error) {
	dst, err = format.Source(src)
	if err != nil {
		dst = src
	}
	return
}

func GoFmtReader(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return GoFmt(b)
}
