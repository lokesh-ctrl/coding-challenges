package Statistics

import "sort"

func InterQuartilesRange(elements []int, frequencies []int) float64 {
	var set []int
	for i := 0; i < len(elements); i++ {
		for j := 0; j < frequencies[i]; j++ {
			set = append(set, elements[i])
		}
	}
	sort.Ints(set)
	quartiles := Quartiles(set)
	interQuartilesRange := quartiles[2] - quartiles[0]
	return float64(interQuartilesRange)
}
