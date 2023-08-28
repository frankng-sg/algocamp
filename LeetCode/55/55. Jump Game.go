package main

import "fmt"

// Time: O(n^2) Space: O(n)
func canJump(nums []int) bool {
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

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // Output: false
	fmt.Println(canJump([]int{3, 2, 1, 1, 4})) // Output: true
}
