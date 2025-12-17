package main

import (
	"reflect"
	"testing"
)

func TestReorderArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "No flags",
			args:     []string{"target"},
			expected: []string{"target"},
		},
		{
			name:     "Flag at start",
			args:     []string{"-i", "target"},
			expected: []string{"-i", "target"},
		},
		{
			name:     "Flag at end",
			args:     []string{"target", "-i"},
			expected: []string{"-i", "target"},
		},
		{
			name:     "Mixed flags",
			args:     []string{"target", "-i", "another code"},
			expected: []string{"-i", "target", "another code"},
		},
		{
			name:     "File flag with argument",
			args:     []string{"target", "-f", "Makefile", "-i"},
			expected: []string{"-f", "Makefile", "-i", "target"},
		},
		{
			name:     "File flag at end",
			args:     []string{"target", "-f", "Makefile"},
			expected: []string{"-f", "Makefile", "target"},
		},
		{
			name:     "Long flag",
			args:     []string{"target", "--print-command"},
			expected: []string{"--print-command", "target"},
		},
		{
			name:     "Complex mix",
			args:     []string{"build", "-i", "-f", "MyMakefile", "--print-command"},
			expected: []string{"-i", "-f", "MyMakefile", "--print-command", "build"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reorderArgs(tt.args)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reorderArgs(%v) = %v, want %v", tt.args, got, tt.expected)
			}
		})
	}
}
