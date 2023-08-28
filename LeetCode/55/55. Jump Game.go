package main

import "fmt"

func canJump(nums []int) bool {
	visited := make([]bool, len(nums))
	visited[0] = true
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			for j := 1; j <= nums[i]; j++ {
				if i+j < len(nums) {
					visited[i+j] = true
				}
			}
		}
	}
	return visited[len(nums)-1]
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4})) // Output: false
	fmt.Println(canJump([]int{3, 2, 1, 1, 4})) // Output: true
}
