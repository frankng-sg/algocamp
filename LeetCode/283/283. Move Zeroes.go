package main

import "fmt"

// Run: O(n) Space: O(1)
func moveZeroes(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	var nums []int

	nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1,3,12,0,0]

	nums = []int{1, 2, 3, 4, 5}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1, 2, 3, 4, 5]

	nums = []int{1}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1]
}
