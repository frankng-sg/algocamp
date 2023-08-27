package main

import "fmt"

// Time: O(n) Space: O(n)
func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[i] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)] = tmp[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // Output: [5 6 7 1 2 3 4]
}
