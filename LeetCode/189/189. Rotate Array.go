package main

import "fmt"

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// Time: O(n) Space: O(1)
func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func main() {
	var nums []int

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 13)
	fmt.Println(nums) // Output: [2 3 4 5 6 7 1]
}
