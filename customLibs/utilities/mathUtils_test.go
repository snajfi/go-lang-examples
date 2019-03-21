package utilities

import (
	"testing"
)

// TestSum is method, which contains test scenario for our utitlity method Sum().
// Test methods has to have name in convention TestXXXX
func TestSum(t *testing.T) {

	cases := []struct{ numberA, numberB, expectedResult int }{
		{5, 3, 8},
		//{1, 1, 3}, // uncoment start of line to see fail message after run
	}

	for _, c := range cases {
		got := Sum(c.numberA, c.numberB)
		if got != c.expectedResult {
			t.Errorf("Sum(%v,%v) == %v, want %v", c.numberA, c.numberB, got, c.expectedResult)
		}
	}
}
