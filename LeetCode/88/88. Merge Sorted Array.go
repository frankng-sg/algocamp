package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	merge(nums, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums) // Output: [1,2,2,3,5,6]
}
