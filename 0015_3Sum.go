package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			s := n + nums[l] + nums[r]

			if s == 0 {
				res = append(res, []int{n, nums[l], nums[r]})
				l++

				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if s > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}
