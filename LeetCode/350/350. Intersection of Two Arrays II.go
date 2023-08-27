package main

import "fmt"

// Time: O(m + n) Space: O(m + n)
func intersect(nums1 []int, nums2 []int) []int {
	var count1, count2 [1001]int
	for _, num := range nums1 {
		count1[num]++
	}
	for _, num := range nums2 {
		count2[num]++
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var output []int
	for i := 0; i <= 1000; i++ {
		k := min(count1[i], count2[i])
		for j := 0; j < k; j++ {
			output = append(output, i)
		}
	}
	return output
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2})) // Output: [2 2]
}
