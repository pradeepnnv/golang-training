package arrays

func Sum(input [5]int) int {
	var sum int
	for i := 0; i < len(input); i++ {
		sum += input[i]
	}
	return sum
}
