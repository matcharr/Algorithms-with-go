package cmd

import "fmt"

func BaseToDec(value string, base int) int {
	res := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		var val int
		fmt.Sscanf(string(value[i]), "%X", &val)
		res += multiplier * val
		multiplier *= base
	}
	return res
	// const charset = "0123456789ABCDEF"
	// res := 0
	// multiplier := 1
	// for i := len(value) - 1; i >= 0; i-- {
	// 	index := -1
	// 	for j, char := range charset {
	// 		if char == rune(value[i]) {
	// 			index = j
	// 			break
	// 		}
	// 	}
	// 	res += index * multiplier
	// 	multiplier *= base
	// }
	// return res
}
