package cli

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestLoginOptions_Complete_URLScheme(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "hostname only",
			args:     []string{"example.com"},
			expected: "https://example.com",
		},
		{
			name:     "hostname with port",
			args:     []string{"example.com:8080"},
			expected: "https://example.com:8080",
		},
		{
			name:     "https scheme present",
			args:     []string{"https://example.com"},
			expected: "https://example.com",
		},
		{
			name:     "http scheme present",
			args:     []string{"http://example.com"},
			expected: "http://example.com",
		},
		{
			name:     "random scheme present",
			args:     []string{"ftp://example.com"},
			expected: "ftp://example.com",
		},
		{
			name:     "empty arg",
			args:     []string{""},
			expected: "",
		},
		{
			name:     "localhost",
			args:     []string{"localhost:8000"},
			expected: "https://localhost:8000",
		},
		{
			name:     "no args",
			args:     []string{},
			expected: "", // No change expected, no error should occur
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			o := DefaultLoginOptions()
			cmd := &cobra.Command{}
			o.Bind(cmd.Flags())

			err := o.Complete(cmd, tc.args)
			require.NoError(t, err)

			if len(tc.args) > 0 {
				require.Equal(t, tc.expected, tc.args[0])
			}
		})
	}
}
