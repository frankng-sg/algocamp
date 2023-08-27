package main

import (
	"fmt"
)

// Time: O(n) Space: O(1)
func singleNumber(nums []int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r ^= nums[i]
	}
	return r
}

func main() {
	fmt.Println(singleNumber([]int{3, 2, 2, 3, 1})) // Output: 1
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2})) // Output: 4
	fmt.Println(singleNumber([]int{4}))             // Output: 4
}
