package main

import (
	"fmt"
)

func trap(height []int) int {
	l := len(height)

	if l <= 2 {
		return 0
	}

	count, lp, rp := 0, 0, 0

	// skip leading 0s
	for i, v := range height {
		if v == 0 {
			continue
		}

		lp = i
		rp = i

		break
	}

	maxIter := l

	for i := l - 1; i >= 0; i-- {
		if height[i] > height[i-1] {
			maxIter = i + 1
			break
		}
	}

	fmt.Println(maxIter, l)

	// iter rp to len
	for rp < maxIter {
		// fmt.Println("lp:", height[lp], "rp:", height[rp])

		// if rp is lower than lp
		if height[rp] < height[lp] {
			tmp := 0
			maxToAdd := 0
			i := rp

			// determine highest next wall
			for i < maxIter {
				// if wall is higher that lp, max = lp
				if height[i] > height[lp] {
					maxToAdd = height[lp]
					break
				}

				// if height is more than max, max = i
				if height[i] > maxToAdd {
					maxToAdd = height[i]
				}

				i++
			}

			fmt.Println("max", maxToAdd)

			// fmt.Println("lp:", height[lp], "rp:", height[rp], maxToAdd)

			// iter rp and count possible fill zones
			for rp < maxIter && height[rp] < height[lp] {
				// fmt.Println("lp:", height[lp], "rp:", height[rp], maxToAdd)
				// fmt.Println("rp:", rp, "adding:", maxToAdd-height[rp], maxToAdd)
				tmp += maxToAdd - height[rp]

				rp++
			}

			// add fill zones to total
			count += tmp
			// increase lp to rp
			lp = rp
		}

		rp++
	}

	return count
}

func main() {
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{3, 0, 0, 0, 1}

	fmt.Println(trap(height))
}
