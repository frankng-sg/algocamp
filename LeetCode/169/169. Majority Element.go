package main

import (
	"fmt"
	"sort"
)

// Time: O(nlogn), Space: O(1)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	count, maxCount, maxVal := 1, 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if count > maxCount {
				maxCount = count
				maxVal = nums[i-1]
			}
			count = 1
		} else {
			count++
		}
	}
	if count > maxCount {
		return nums[len(nums)-1]
	}
	return maxVal
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             // Output: 3
	fmt.Println(majorityElement([]int{1, 0, 3, 2, 2, 2, 2})) // Output: 2
	fmt.Println(majorityElement([]int{1}))                   // Output: 1
}
