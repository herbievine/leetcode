package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, v := range nums {
		_, ok := m[v]

		if ok {
			return true
		} else {
			m[v] = true
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(nums))
}
