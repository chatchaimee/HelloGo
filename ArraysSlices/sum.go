package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// var sums []int
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		// numbers[1] Result is x
		// numbers[1:] is mean return all value index 1 - n to array, result is [x...n]
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
