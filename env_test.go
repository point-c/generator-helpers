package generator_helpers

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTestOutputFilename(t *testing.T) {
	require.Equal(t, "foo_test.go", TestOutputFilename("foo.go"))
	require.Equal(t, "foo.sh_test.go", TestOutputFilename("foo.sh"))
}

func Test_OutputFilename(t *testing.T) {
	for _, c := range []struct {
		name, input, expected string
	}{
		{
			name:     "with ext",
			input:    "test.go",
			expected: "test_generated.go",
		},
		{
			name:     "without ext",
			input:    "test",
			expected: "test_generated.go",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.expected, OutputFilename(c.input))
		})
	}
}

func TestEnvOrDefault(t *testing.T) {
	type args struct {
		key string
		def string
	}
	type test struct {
		name string
		args args
		want string
		env  map[string]string
	}

	tests := []test{
		{
			name: "default",
			args: args{
				key: uuid.New().String(),
				def: "test",
			},
			want: "test",
		},
		func(key string) test {
			return test{
				name: "env",
				args: args{key: key},
				want: "test",
				env:  map[string]string{key: "test"},
			}
		}(uuid.New().String()),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(tt.args.key, "")
			if tt.env != nil {
				for k, v := range tt.env {
					t.Setenv(k, v)
				}
			}

			require.Equal(t, tt.want, EnvOrDefault(tt.args.key, tt.args.def))
		})
	}
}

func TestEnvOrDefaultInt(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Setenv("test", "123")
		v := EnvOrDefaultInt("test", 0)
		require.Equal(t, 123, v)
	})
	t.Run("invalid", func(t *testing.T) {
		t.Setenv("test", "a")
		v := EnvOrDefaultInt("test", 123)
		require.Equal(t, 123, v)
	})
}

func TestIgnoreForCoverage(t *testing.T) {
	// These will only return default unless called by go generate
	_ = EnvGoFile()
	_ = EnvGoPackage()
	_ = EnvGoArch()
	_ = EnvGoOS()
	_ = EnvGoLine()
	_ = EnvGoRoot()
	_ = EnvDollar()
	_ = EnvPath()
}

func TestIfStringEmptyUseDefault(t *testing.T) {
	type args struct {
		s   string
		def string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "non-empty",
			args: args{
				s:   "foo",
				def: "bar",
			},
			want: "foo",
		},
		{
			name: "empty",
			args: args{
				s:   "",
				def: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfStringEmptyUseDefault(tt.args.s, tt.args.def); got != tt.want {
				t.Errorf("IfStringEmptyUseDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
