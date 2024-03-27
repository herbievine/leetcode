package main

import "fmt"

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	output := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = prefix[i-1] * nums[i]
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			suffix[i] = nums[i]
		} else {
			suffix[i] = suffix[i+1] * nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			output[i] = suffix[i+1]
		} else if i == len(nums)-1 {
			output[i] = prefix[i-1]
		} else {
			output[i] = prefix[i-1] * suffix[i+1]
		}
	}

	return output
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(productExceptSelf(nums))
}
