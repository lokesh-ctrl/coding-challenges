package challenges

func SumEvenAfterQueries(A []int, queries [4][2]int) []int {
	var results []int
	var evensum int
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			evensum = evensum + A[i]
		}
	}
	for i := 0; i < len(A); i++ {
		if A[queries[i][1]]%2 == 0 {
			evensum = evensum - A[queries[i][1]]
		}
		A[queries[i][1]] = A[queries[i][1]] + queries[i][0]
		if A[queries[i][1]]%2 == 0 {
			evensum = evensum + A[queries[i][1]]
		}
		results = append(results, evensum)
	}
	return results
}
