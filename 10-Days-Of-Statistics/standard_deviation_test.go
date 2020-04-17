package Statistics

import (
	"testing"
)

func TestStandardDeviation(t *testing.T) {
	var tests = []struct {
		input    []int
		expected float64
	}{
		{[]int{10, 40, 30, 50, 20}, 14.1},
	}

	for _, test := range tests {
		if output := StandardDeviation(test.input); test.expected != output {
			t.Error("Test Failed: inputted, {} expected, recieved: {}", test.expected, output)
		}
	}
}
