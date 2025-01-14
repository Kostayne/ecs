package utils

import (
	"testing"

	. "github.com/kostayne/ecs/utils"
)

func TestFastRemoveI(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		i        int
		expected []int
	}{
		{"remove first element", []int{1, 2, 3}, 0, []int{3, 2}},
		{"remove middle element", []int{1, 2, 3}, 1, []int{1, 3}},
		{"remove last element", []int{1, 2, 3}, 2, []int{1, 2}},
		{"empty array", []int{}, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FastRemoveI(tt.arr, tt.i)

			if !sliceEqual(actual, tt.expected) {
				t.Errorf("FastRemoveI(%v, %d) = %v, want %v", tt.arr, tt.i, actual, tt.expected)
			}
		})
	}

	// test index out of range
	t.Run("index out of range", func(t *testing.T) {
		defer func() {
			if r := recover(); r != "Index out of range!" {
				t.Errorf("FastRemoveI([]int{1, 2, 3}, 3) did not panic with 'Index out of range!'")
			}
		}()

		FastRemoveI([]int{1, 2, 3}, 3)
	})
}

func TestFastRemove(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		item     int
		expected []int
	}{
		{"remove first element", []int{1, 2, 3}, 1, []int{3, 2}},
		{"remove middle element", []int{1, 2, 3}, 2, []int{1, 3}},
		{"remove last element", []int{1, 2, 3}, 3, []int{1, 2}},
		{"item not found", []int{1, 2, 3}, 4, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FastRemove(tt.arr, tt.item)

			if !sliceEqual(actual, tt.expected) {
				t.Errorf("FastRemove(%v, %d) = %v, want %v", tt.arr, tt.item, actual, tt.expected)
			}
		})
	}
}

func TestShiftRemoveI(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		i        int
		expected []int
	}{
		{"remove first element", []int{1, 2, 3}, 0, []int{2, 3}},
		{"remove middle element", []int{1, 2, 3}, 1, []int{1, 3}},
		{"remove last element", []int{1, 2, 3}, 2, []int{1, 2}},
		{"empty array", []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ShiftRemoveI(tt.arr, tt.i)

			if !sliceEqual(actual, tt.expected) {
				t.Errorf("ShiftRemoveI(%v, %d) = %v, want %v", tt.arr, tt.i, actual, tt.expected)
			}
		})
	}

	// test index out of range
	t.Run("index out of range", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("ShiftRemoveI([]int{1, 2, 3}, 3) did not panic")
			}
		}()

		ShiftRemoveI([]int{1, 2, 3}, 3)
	})
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
