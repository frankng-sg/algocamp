package main

import (
	"fmt"
	"sort"
)

// Time: O(m + n) Space: O(1)
func intersect(nums1 []int, nums2 []int) []int {
	var count [1001]int
	for _, num := range nums1 {
		count[num]++
	}
	output := make([]int, 0, len(nums1))
	for _, num := range nums2 {
		if count[num] > 0 {
			count[num]--
			output = append(output, num)
		}
	}
	return output
}

// Time: O(Max(nlogn, m + n)) Space: O(1)
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	output := make([]int, 0, len(nums1))
	var i, j int
	for ; i < len(nums1); i++ {
		for ; j < len(nums2) && nums2[j] < nums1[i]; j++ {
		}
		if j >= len(nums2) {
			break
		}
		if nums2[j] == nums1[i] {
			output = append(output, nums2[j])
			j++
		}
	}
	return output
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))  // Output: [2 2]
	fmt.Println(intersect2([]int{1, 2, 2, 1}, []int{2, 2})) // Output: [2 2]
}
