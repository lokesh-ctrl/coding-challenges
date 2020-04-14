package Statistics

func WeightedMean(elements []int, weights []int) float64 {
	WeightedSum := 0
	WeightSum := 0
	for i := 0; i < len(elements); i++ {
		WeightedSum += elements[i] * weights[i]
		WeightSum += weights[i]
	}
	FWeightedSum := float64(WeightedSum)
	FWeightSum := float64(WeightSum)

	return FWeightedSum / FWeightSum
}
