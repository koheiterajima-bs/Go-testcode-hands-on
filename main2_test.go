package main

import "testing"

func TestHelloTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test with Kohei", "Kohei", "Hello, Kohei!"},
		{"Test with Alice", "Alice", "Hello, Alice!"},
		{"Test with Bob", "Bob", "Hello, Bob!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.input)
			if got != tt.want {
				t.Errorf("Hello(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
