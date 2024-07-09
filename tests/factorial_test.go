package tests

import (
	"testing"

	"paulobraga.com/study/lib"
)

func TestFactRecursive(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{0, 1},   // Factorial of 0
		{5, 120}, // Factorial of a positive number
		{-5, -1}, // Factorial of a negative number
	}

	for _, test := range tests {
		result := lib.FactRecursive(test.num)
		if result != test.expected {
			t.Errorf("Expected %d for input %d, but got %d", test.expected, test.num, result)
		}
	}
}
