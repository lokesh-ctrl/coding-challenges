package Statistics

import (
	"sort"
)

func Quartiles(list []int) []int {
	var Q2, Q1, Q3 float64
	var results []int
	sort.Ints(list)
	Q2 = Median(list, len(list))
	if len(list)%2 == 0 {
		lowerhalf := list[0 : len(list)/2 : len(list)/2]
		upperhalf := list[len(list)/2 : len(list) : len(list)]
		Q1 = Median(lowerhalf, len(lowerhalf))
		Q3 = Median(upperhalf, len(upperhalf))
		results = append(results, int(Q1))
		results = append(results, int(Q2))
		results = append(results, int(Q3))
	} else {
		lowerhalf := list[0 : len(list)/2 : len(list)/2]
		upperhalf := list[len(list)/2+1 : len(list) : len(list)]
		Q1 = Median(lowerhalf, len(lowerhalf))
		Q3 = Median(upperhalf, len(upperhalf))
		results = append(results, int(Q1))
		results = append(results, int(Q2))
		results = append(results, int(Q3))
	}
	return results
}

func Median(list []int, n int) float64 {
	if (n % 2) != 0 {
		counter := (n / 2)
		return float64(list[counter])
	}
	sum := list[n/2] + list[(n/2)-1]
	sumFloat := float64(sum)
	return (sumFloat / float64(2))
}
