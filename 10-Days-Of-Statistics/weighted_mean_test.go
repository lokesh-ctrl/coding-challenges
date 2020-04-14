package Statistics

import "testing"

func TestWeightedMean(t *testing.T) {
	weights := make([]int, 0)
	elements := make([]int, 0)

	elements = append(elements, 10, 40, 30, 50, 20)
	weights = append(weights, 1, 2, 3, 4, 5)

	ans := WeightedMean(elements, weights)

	if ans != 32.0 {
		t.Errorf("Expecting %f but got %f", 32.0, ans)
	}
}
