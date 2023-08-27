package main

import "fmt"

// Time: O(n)
// Mem: O(1)
func removeDuplicates(nums []int) int {
	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{3, 2, 2, 3}))               // Output: 2
	fmt.Println(removeDuplicates([]int{10, 9, 8, 7, 8, 9, 10, 3})) // Output: 5
}
