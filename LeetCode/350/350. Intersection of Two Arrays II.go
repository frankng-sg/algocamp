package main

import "fmt"

// Time: O(m + n) Space: O(1)
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
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

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2})) // Output: [2 2]
}
