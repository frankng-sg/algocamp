package main

import (
	"fmt"
	"sort"
)

func partition(nums []int, l, r int) int {
	for i := l; i < r; i++ {
		if nums[i] >= nums[r] {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}
	nums[l], nums[r] = nums[r], nums[l]
	return l
}

func qselect(nums []int, l, r, k int) int {
	pivot := partition(nums, l, r)
	if k > pivot {
		return qselect(nums, pivot+1, r, k)
	} else if k < pivot {
		return qselect(nums, l, pivot-1, k)
	}
	return nums[k]
}

// Algorithm: Quick Select
// Time: O(n^2), Space: O(1)
func findKthLargest1(nums []int, k int) int {
	return qselect(nums, 0, len(nums)-1, k-1)
}

// Algorithm: Quick Sort
// Time: O(nlogn), Space: O(1)
func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return nums[k-1]
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // Output: 5
	fmt.Println(findKthLargest([]int{99, 99}, 1))                    // Output: 99
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // Output: 4
	fmt.Println(findKthLargest([]int{1}, 1))                         // Output: 1
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // Output: 5
}
