package main

import (
	"fmt"
)

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	mArea, width := 0, right-left
	for left < right {
		area := lowerInt(height[left], height[right]) * width
		if area > mArea {
			mArea = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
		width--
	}
	return mArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))  // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                       // Output: 1
	fmt.Println(maxArea([]int{1, 1, 1, 100, 101, 1, 1, 1})) // Output: 100
}
