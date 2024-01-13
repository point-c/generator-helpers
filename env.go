package generator_helpers

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func OutputFilename(original string) string {
	return strings.TrimSuffix(original, ".go") + "_generated.go"
}

func TestOutputFilename(original string) string {
	return strings.TrimSuffix(original, ".go") + "_generated_test.go"
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

func EnvGoFile() string    { return EnvOrDefault(EnvKeyGoFile, EnvDefaultGoFile) }
func EnvGoPackage() string { return EnvOrDefault(EnvKeyGoPackage, EnvDefaultGoPackage) }
func EnvGoArch() string    { return EnvOrDefault(EnvKeyGoArch, EnvDefaultGoArch) }
func EnvGoOS() string      { return EnvOrDefault(EnvKeyGoOS, EnvDefaultGoOS) }
func EnvGoLine() int       { return EnvOrDefaultInt(EnvKeyGoLine, EnvDefaultGoLine) }
func EnvGoRoot() string    { return EnvOrDefault(EnvKeyGoRoot, EnvDefaultGoRoot) }
func EnvDollar() string    { return EnvOrDefault(EnvKeyDollar, EnvDefaultDollar) }
func EnvPath() string      { return EnvOrDefault(EnvKeyPath, EnvDefaultPath) }

func EnvOrDefault(key, def string) string {
	return IfStringEmptyUseDefault(os.Getenv(key), def)
}

func EnvOrDefaultInt(key string, def int) int {
	v := EnvOrDefault(key, strconv.Itoa(def))
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	return def
}

func IfStringEmptyUseDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
