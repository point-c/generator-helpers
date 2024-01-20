package generator_helpers

import (
	"log/slog"
	"os"
)

var exit = os.Exit

// Must takes a value and an error. If the error is not nil, it logs the error and exits the program.
// Otherwise, it returns the value.
func Must[T any](t T, err error) T {
	if err != nil {
		slog.Error("failed to get value", "error", err)
		exit(1)
	}
	return t
}

// Check takes an error as an argument. If the error is not nil, it logs the error and exits the program.
func Check(err error) {
	if err != nil {
		slog.Error("operation failed", "error", err)
		exit(1)
	}
}
