package cmd

func Factor(primes []int, number int) []int {
	var res []int

	if number > 1 {
		res = append(res, number)
	}

	for _, prime := range primes {
		for number%prime == 0 {
			res = append(res, prime)
			number /= prime
		}
	}
	return res
}
