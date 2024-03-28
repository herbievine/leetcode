package main

import (
	"fmt"
)

func isAlnum(c rune) bool {
	if (c >= 65 && c <= 90) || (c >= 48 && c <= 57) || (c >= 97 && c <= 122) {
		return true
	}

	return false
}

func isPalindrome(s string) bool {
	v := make([]rune, len(s))

	rp := 0
	for _, c := range s {
		if isAlnum(c) {
			if c >= 65 && c <= 90 {
				v[rp] = c + 32
			} else {
				v[rp] = c
			}

			rp++
		}
	}

	lp := 0

	if rp > 0 {
		rp--
	}

	for lp <= rp {
		if v[lp] != v[rp] {
			return false
		}

		lp++
		rp--
	}

	return true
}

func main() {
	str := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(str))
}
