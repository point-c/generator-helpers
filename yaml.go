package generator_helpers

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

// UnmarshalYAML reads YAML data from a file and unmarshals it into a value.
// It returns the unmarshaled value and any error encountered.
func UnmarshalYAML[T any](filename string) (v T, err error) {
	var f *os.File
	f, err = os.Open(filename)
	if err != nil {
		return v, err
	}
	defer func() { errors.Join(err, f.Close()) }()

	if err = yaml.NewDecoder(f).Decode(&v); err != nil {
		return v, err
	}
	return
}
