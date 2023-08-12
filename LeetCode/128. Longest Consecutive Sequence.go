package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	nlen := len(nums)
	if nlen <= 0 {
		return 0
	}
	sort.Ints(nums)
	max := 1
	count := 1
	for i := 1; i < nlen; i++ {
		if nums[i]-nums[i-1] == 1 {
			count++
		} else if nums[i]-nums[i-1] > 1 {
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{9, 8, 7, 6, 5, 4}))             // Output: 6
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         // Output: 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // Output: 9
	fmt.Println(longestConsecutive([]int{0, 1, 1, 2}))                   // Output: 3
}
