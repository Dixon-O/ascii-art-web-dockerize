package artprinter

import (
	"strings"
	"testing"
)

// TestCheckOnlyNewLines() checks if the function CheckOnlyNewLines() works correctly
func TestCheckOnlyNewLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"single newline", "\n", true},
		{"two newline chars", "\n\n", true},
		{"n char in between two newline chars", "\nn\n", false},
		{"triple newline chars", "\n\n\n", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckOnlyNewLines(strings.Split(tt.input, "\n"))
			if result != tt.expected {
				t.Errorf("CheckOnlyNewLines() with input %q returns %t, want %t",
					tt.input, result, tt.expected)
			}
		})
	}
}
