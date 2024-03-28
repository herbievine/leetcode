package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	lp := 0
	rp := len(numbers) - 1

	for true {
		if numbers[lp]+numbers[rp] == target {
			return []int{lp + 1, rp + 1}
		}

		if numbers[lp]+numbers[rp] > target {
			rp--
		} else {
			lp++
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum(nums, 9))
}
