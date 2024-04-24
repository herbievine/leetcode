package main

import (
	"fmt"
)

func trap(height []int) int {
	l := len(height)

	if l <= 2 {
		return 0
	}

	total, lp, i, rp := 0, 0, 2, 2

	for i < l {
		if height[i] > height[rp] {
			rp = i
		}

		i++
	}

	i = 1

	for i < l-1 {
		res := height[lp] - height[i]

		if height[lp] > height[rp] {
			res = height[rp] - height[i]
		}

		if res > 0 {
			total += res
		}

		if height[i] > height[lp] {
			lp = i
		}

		if i == rp-1 {
			rp++
			j := rp

			for j < l {
				if height[j] > height[rp] {
					rp = j
				}

				j++
			}
		}

		i++
	}

	return total
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(height))
}
