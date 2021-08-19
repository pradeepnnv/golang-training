package arrays

func Sum(input []int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}
