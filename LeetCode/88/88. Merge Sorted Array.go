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

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
	for ; j >= 0; j, k = j-1, k-1 {
		nums1[k] = nums2[j]
	}
}

func main() {
	nums := []int{1, 2, 3, 0, 0, 0}
	merge2(nums, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums) // Output: [1,2,2,3,5,6]
}
