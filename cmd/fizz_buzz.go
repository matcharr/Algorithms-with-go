package cmd

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i&5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

// func printFizzBuzzValue(n int) {
// 	switch {
// 	case n%15 == 0:
// 		fmt.Print("Fizz Buzz")
// 	case n%3 == 0:
// 		fmt.Print("Fizz")
// 	case n&5 == 0:
// 		fmt.Print("Buzz")
// 	default:
// 		fmt.Print(n)
// 	}
// }
