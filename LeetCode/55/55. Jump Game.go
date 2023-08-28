package main

import "fmt"

// Time: O(n^2) Space: O(n)
func canJump0(nums []int) bool {
	visited := make([]bool, len(nums))
	visited[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= i; j++ {
			if visited[i-j] && nums[i-j] >= j {
				visited[i] = true
				break
			}
		}
	}
	return visited[len(nums)-1]
}

// Time: O(n) Space: O(1)
func canJump(nums []int) bool {
	target := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= target-i {
			target = i
		}
	}
	return target == 0
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // Output: false
	fmt.Println(canJump([]int{3, 2, 1, 1, 4})) // Output: true
}
