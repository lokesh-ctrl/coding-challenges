package Statistics

import "math"

func StandardDeviation(list []int) float64 {
	mean := Mean(list)
	var squaredDifferenceSum float64
	for _, item := range list {
		squaredDifferenceSum += ((float64(item) - mean) * (float64(item) - mean))
	}
	squaredDifferenceSumDividedByNumberList := squaredDifferenceSum / float64(len(list))
	standardDeviation := math.Sqrt(squaredDifferenceSumDividedByNumberList)
	//Rounding to one decimal
	standardDeviation = math.Round(standardDeviation*10) / 10
	return standardDeviation
}

func Mean(list []int) float64 {
	var sum int
	for _, item := range list {
		sum += item
	}
	return float64(sum) / float64(len(list))
}
