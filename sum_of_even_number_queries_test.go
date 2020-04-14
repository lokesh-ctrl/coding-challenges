package challenges

import (
	"reflect"
	"testing"
)

func TestSumEvenAfterQueries(t *testing.T) {
	A := make([]int, 0)
	A = append(A, 1, 2, 3, 4)
	queries := [4][2]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	result := []int{8, 6, 2, 4}
	total := SumEvenAfterQueries(A, queries)
	if !reflect.DeepEqual(result, total) {
		t.Errorf("Sum was incorrect")
	}
}
