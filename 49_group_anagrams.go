package main

import (
	"fmt"
)

func wordKey(s string) [26]uint8 {
	var key [26]uint8

	for _, v := range s {
		key[v-'a']++
	}

	return key
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]uint8][]string)

	for _, s := range strs {
		key := wordKey(s)
		m[key] = append(m[key], s)
	}

	r, i := make([][]string, len(m)), 0

	for _, v := range m {
		r[i] = v
		i++
	}

	return r
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
