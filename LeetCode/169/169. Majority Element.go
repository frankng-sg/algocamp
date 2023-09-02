package main

import (
	"fmt"
	"sort"
)

// Boyer-Moore Majority Voting Algorithm
// Time: O(n), Space: O(1)
func majorityElement(nums []int) int {
	vote, candidate := 1, nums[0]
	for _, v := range nums {
		if v == candidate {
			vote++
		} else {
			vote--
		}
		if vote == 0 {
			candidate = v
			vote = 1
		}
	}
	return candidate
}

// Time: O(n), Space: O(n)
func majorityElement2(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	for _, v := range nums {
		if count[v] > len(nums)>>1 {
			return v
		}
	}
	return nums[0]
}

// Time: O(nlogn), Space: O(1)
func majorityElement1(nums []int) int {
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
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))       // Output: 2
	fmt.Println(majorityElement([]int{3, 2, 3}))                   // Output: 3
	fmt.Println(majorityElement([]int{1, 0, 3, 2, 2, 2, 2}))       // Output: 2
	fmt.Println(majorityElement([]int{1}))                         // Output: 1
	fmt.Println(majorityElement([]int{1, 0, 2, 0, 3, 0, 4, 0, 0})) // Output: 0
	fmt.Println(majorityElement([]int{1, 1, 1, 2, 2, 2, 2}))       // Output: 2
}
