package cmd

func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}
