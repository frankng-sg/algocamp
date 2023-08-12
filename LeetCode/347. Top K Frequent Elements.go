package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v] += 1
	}
	var keys []int
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	return keys[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // Output: [1, 2]
	fmt.Println(topKFrequent([]int{1}, 1))                // Output: [1]
	fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))       // Output: [0]
}
