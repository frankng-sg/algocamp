package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	k := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-1] || nums[i] != nums[k-1] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}

func main() {
	var nums []int
	var k int

	nums = []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	k = removeDuplicates(nums)
	fmt.Println(nums[:k]) // Output [1 1 2 2 3 3]

	nums = []int{1, 2, 2, 2, 2, 2, 3, 3, 3}
	k = removeDuplicates(nums)
	fmt.Println(nums[:k]) // Output [1 2 2 3 3]

	nums = []int{1}
	k = removeDuplicates(nums)
	fmt.Println(nums[:k]) // Output [1 2 2 3 3]
}
