package Statistics

import (
	"reflect"
	"testing"
)

func TestQuartiles(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{3, 7, 8, 5, 12, 14, 21, 13, 18}, []int{6, 12, 16}},
		{[]int{3, 7, 8, 5, 12, 14, 21, 15, 18, 14}, []int{7, 13, 15}},
	}

	for _, test := range tests {
		if output := Quartiles(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error("Test Failed: inputted, {} expected, recieved: {}", test.expected, output)
		}
	}
}

func TestMedian(t *testing.T) {
	var tests = []struct {
		input    []int
		expected float64
	}{
		{[]int{3, 7, 8, 5, 12, 14, 21, 13, 18}, 12},
		{[]int{1, 3, 4, 5, 7, 8, 9, 11}, 6},
	}

	for _, test := range tests {
		if output := Median(test.input, len(test.input)); output != test.expected {
			t.Error("Test Failed: inputted, {} expected, recieved: {}", test.expected, output)
		}
	}
}
