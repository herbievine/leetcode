package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	r := 0

	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		if _, ok := m[v-1]; ok {
			continue
		}

		l := 0

		for m[v+l] {
			if _, ok := m[v+l]; ok {
				l++
			}
		}

		if l > r {
			r = l
		}

		if l > len(nums)/2 {
			break
		}
	}

	return r
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(longestConsecutive(nums))
}
