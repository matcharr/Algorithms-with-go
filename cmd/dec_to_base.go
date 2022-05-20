package cmd

import "strings"

func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var sb strings.Builder
	for dec > 0 {
		rem := dec % base
		sb.WriteByte(charset[rem])
		dec = dec / base
	}
	return Reverse(sb.String())
}
