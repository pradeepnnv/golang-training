package arrays

// Sum returns sum of all integers in the given slice.
func Sum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}

// SumAll returns a slice of integers representing sum of all integers in each input slice.
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sum := 0
		for _, i := range v {
			sum += i
		}
		sums = append(sums, sum)
	}
	return
}

// SumAllTails returns a slice of ints representing sum of all integers (except the first one) in each input slice.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(v[1:]))
		}
	}
	return
}
