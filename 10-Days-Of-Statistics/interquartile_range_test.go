package Statistics

import (
	"testing"
)

func TestInterQuartilesRange(t *testing.T) {
	var tests = []struct {
		input1   []int
		input2   []int
		expected float64
	}{
		{[]int{6, 12, 8, 10, 20, 16}, []int{5, 4, 3, 2, 1, 5}, 9.0},
	}

	for _, test := range tests {
		if output := InterQuartilesRange(test.input1, test.input2); test.expected != output {
			t.Error("Test Failed: inputted, {} expected, recieved: {}", test.expected, output)
		}
	}
}
