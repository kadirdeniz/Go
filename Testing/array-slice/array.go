package array

func Sum(arr []int) int {
	response := 0
	for _, number := range arr {
		response += number
	}
	return response
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
