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
	for i := 1; i < len(s); i++ {
		if (s[i] - '0') < (s[i-1] - '0') {
			return false
		}
		if s[i] == s[i-1] {
			double = true
		}
	}
	return double
}
