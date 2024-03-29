package main

import (
	"fmt"
)

func calcArea(x, y, xh, yh int) int {
	h := xh

	if xh > yh {
		h = yh
	}

	return (y - x) * h
}

func maxArea(height []int) int {
	lp, rp := 0, len(height)-1
	max := calcArea(lp, rp, height[lp], height[rp])

	if len(height) == 2 {
		return max
	}

	for lp <= rp {
		t := calcArea(lp, rp, height[lp], height[rp])

		if t > max {
			max = t
		}

		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
	}

	return max
}

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(nums))
}
