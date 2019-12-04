package main

import (
	"fmt"
	"strconv"
)

func main() {
	correct := 0

	for i := 347312; i <= 805915; i++ {
		s := strconv.Itoa(i)
		if satisfiesCriteria(s) {
			correct++
		}
	}

	fmt.Println(correct)
}

func satisfiesCriteria(s string) bool {
	double := false

	for i := 0; i < len(s)-1; i++ {
		if (s[i] - '0') > (s[i+1] - '0') {
			return false
		}

		if (i == 0 || s[i] != s[i-1]) && s[i] == s[i+1] && (i+2 >= len(s) || s[i] != s[i+2]) {
			double = true
		}
	}
	return double
}
