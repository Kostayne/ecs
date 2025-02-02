package utils_test

import (
	"testing"

	. "github.com/kostayne/ecs/v2/utils"
)

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		arr      []string
		value    string
		expected int
	}{
		{"found at start", []string{"a", "b", "c"}, "a", 0},
		{"found in middle", []string{"a", "b", "c"}, "b", 1},
		{"found at end", []string{"a", "b", "c"}, "c", 2},
		{"not found", []string{"a", "b", "c"}, "d", -1},
		{"empty array", []string{}, "a", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IndexOf(tt.arr, tt.value)

			if actual != tt.expected {
				t.Errorf("IndexOf(%v, %s) = %d, want %d", tt.arr, tt.value, actual, tt.expected)
			}
		})
	}
}
