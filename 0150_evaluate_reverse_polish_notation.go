package main

import (
	"fmt"
	"strconv"
)

type Stack []int

func (s *Stack) Push(str int) {
	*s = append(*s, str)
}

func (s *Stack) Pop() int {
	l := len(*s)

	if l == 0 {
		panic("impossible")
	}

	val := (*s)[l-1]
	*s = (*s)[:l-1]

	return val
}

func evalRPN(tokens []string) int {
	stack := Stack{}

	for _, token := range tokens {
		val, err := strconv.ParseInt(token, 10, 16)
		if err != nil {
			a := stack.Pop()
			b := stack.Pop()

			if token == "+" {
				stack.Push(b + a)
			} else if token == "-" {
				stack.Push(b - a)
			} else if token == "*" {
				stack.Push(b * a)
			} else if token == "/" {
				stack.Push(b / a)
			}

			continue
		}

		stack.Push(int(val))
	}

	return stack.Pop()
}

func main() {
	nums := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	fmt.Println(evalRPN(nums))
}
