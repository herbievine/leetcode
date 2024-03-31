package main

import "fmt"

func search(nums []int, target int) int {
	l, h := 0, len(nums)-1

	for l <= h {
		m := l + (h-l)/2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9))
}
