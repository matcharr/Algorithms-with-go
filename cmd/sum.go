package cmd

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}

// Sum(3,4,5) =>
// 	3 + Sum(4,5) =>
// 		4 + Sum(5) =>
// 			5 + Sum() =>
// 				0
