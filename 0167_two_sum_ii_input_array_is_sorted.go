package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	l := len(numbers)

	if l == 2 {
		return []int{1, 2}
	}

	for i, iv := range numbers {
		for j, jv := range numbers {
			if i == j {
				continue
			}

			if iv+jv == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))
}
