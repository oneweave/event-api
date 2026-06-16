package lib

import "testing"

func TestWithTrailingSlash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"already has slash", "path/", "path/"},
		{"missing slash", "path", "path/"},
		{"root slash", "/", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := withTrailingSlash(tt.input)
			if result != tt.expected {
				t.Fatalf("withTrailingSlash(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestWithLeadingSlash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"already has slash", "/path", "/path"},
		{"missing slash", "path", "/path"},
		{"single slash", "/", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := withLeadingSlash(tt.input)
			if result != tt.expected {
				t.Fatalf("withLeadingSlash(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
