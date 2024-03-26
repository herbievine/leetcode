package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	bucket := make(map[int][]int)

	for k, v := range m {
		if _, ok := bucket[v]; !ok {
			bucket[v] = []int{}
		}

		bucket[v] = append(bucket[v], k)
	}

	res := []int{}

	for i := len(nums); i > 0; i-- {
		if v, ok := bucket[i]; ok {
			res = append(res, v...)
		}
	}

	return res[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(topKFrequent(nums, k))

	nums = []int{-1, -1}
	k = 1

	fmt.Println(topKFrequent(nums, k))

	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	k = 3

	fmt.Println(topKFrequent(nums, k))
}
