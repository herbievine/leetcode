package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, len(s))

	n := 0

	for _, c := range s {
		if c == '(' {
			stack = append(stack, '(')
			n++
		} else if c == '[' {
			stack = append(stack, '[')
			n++
		} else if c == '{' {
			stack = append(stack, '{')
			n++
		} else if c == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			n--
		} else if c == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
			n--
		} else if c == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
			n--
		} else {
			return false
		}
	}

	return n == 0
}

func main() {
	s := "(("

	fmt.Println(isValid(s))
}
