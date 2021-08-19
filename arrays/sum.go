package arrays

func Sum(input [5]int) int {
	var sum int
	for _, v := range input {
		sum += v
	}
	return sum
}
