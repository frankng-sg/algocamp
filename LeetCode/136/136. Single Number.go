package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn) Space: O(1)
func singleNumber(nums []int) int {
	sort.Ints(nums)
	var i int
	for i = 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	if i == len(nums)-1 {
		return nums[i]
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{3, 2, 2, 3, 1})) // Output: 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // Output: 4
	fmt.Println(singleNumber([]int{4}))             // Output: 4
}
