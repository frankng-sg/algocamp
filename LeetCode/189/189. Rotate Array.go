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

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0

	for start := 0; count < n; start++ {
		current, tmp1 := start, nums[start]
		for {
			next := (current + k) % n
			tmp2 := nums[next]
			nums[next] = tmp1
			tmp1 = tmp2
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}

func main() {
	var nums []int

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 13)
	fmt.Println(nums) // Output: [2 3 4 5 6 7 1]

	nums = []int{-1, -100, 3, 99}
	rotate2(nums, 2)
	fmt.Println(nums) // Output: [3,99,-1,-100]

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(nums, 3)
	fmt.Println(nums) // Output: [5,6,7,1,2,3,4]
}
