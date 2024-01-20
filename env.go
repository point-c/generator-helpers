package generator_helpers

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

// OutputFilename generates a filename for the output file by appending "_generated" to the original file name.
func OutputFilename(original string) string {
	return strings.TrimSuffix(original, ".go") + "_generated.go"
}

// TestOutputFilename generates a filename for the test output file by appending "_test" to the original file name.
func TestOutputFilename(original string) string {
	return strings.TrimSuffix(original, ".go") + "_test.go"
}

var (
	EnvDefaultGoPackage = "generator_helpers"
	EnvDefaultGoFile    = ""
	EnvDefaultGoArch    = runtime.GOARCH
	EnvDefaultGoOS      = runtime.GOOS
	EnvDefaultGoLine    = 0
	EnvDefaultDollar    = "$"
	EnvDefaultPath      = os.Getenv(EnvKeyPath)
	EnvDefaultGoRoot    = os.Getenv(EnvKeyGoRoot)
)

func init() {
	_, EnvDefaultGoFile, EnvDefaultGoLine, _ = runtime.Caller(0)
}

const (
	EnvKeyGoArch    = "GOARCH"
	EnvKeyGoOS      = "GOOS"
	EnvKeyGoFile    = "GOFILE"
	EnvKeyGoLine    = "GOLINE"
	EnvKeyGoPackage = "GOPACKAGE"
	EnvKeyGoRoot    = "GOROOT"
	EnvKeyDollar    = "DOLLAR"
	EnvKeyPath      = "PATH"
)

// EnvGoFile returns the Go file name from environment or default.
func EnvGoFile() string { return EnvOrDefault(EnvKeyGoFile, EnvDefaultGoFile) }

// EnvGoPackage returns the Go package name from environment or default.
func EnvGoPackage() string { return EnvOrDefault(EnvKeyGoPackage, EnvDefaultGoPackage) }

// EnvGoArch returns the Go architecture from environment or default.
func EnvGoArch() string { return EnvOrDefault(EnvKeyGoArch, EnvDefaultGoArch) }

// EnvGoOS returns the Go operating system from environment or default.
func EnvGoOS() string { return EnvOrDefault(EnvKeyGoOS, EnvDefaultGoOS) }

// EnvGoLine returns the Go file line number from environment or default.
func EnvGoLine() int { return EnvOrDefaultInt(EnvKeyGoLine, EnvDefaultGoLine) }

// EnvGoRoot returns the Go file root path from environment or default.
func EnvGoRoot() string { return EnvOrDefault(EnvKeyGoRoot, EnvDefaultGoRoot) }

// EnvDollar returns the dollar sign from environment or default.
func EnvDollar() string { return EnvOrDefault(EnvKeyDollar, EnvDefaultDollar) }

// EnvPath returns the system path from environment or default.
func EnvPath() string { return EnvOrDefault(EnvKeyPath, EnvDefaultPath) }

// EnvOrDefault returns the environment variable's value or default if empty.
func EnvOrDefault(key, def string) string {
	return IfStringEmptyUseDefault(os.Getenv(key), def)
}

// EnvOrDefaultInt returns the environment variable's integer value or default if empty.
// If there is an error parsing the environment variable. The default will be returned.
func EnvOrDefaultInt(key string, def int) int {
	v := EnvOrDefault(key, strconv.Itoa(def))
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	return def
}

// IfStringEmptyUseDefault returns the default value if the string is empty.
func IfStringEmptyUseDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
