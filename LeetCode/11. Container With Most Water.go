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
	left := 0
	right := len(height) - 1
	mArea := 0
	for left < right {
		lower := lowerInt(height[left], height[right])
		area := lower * (right - left)
		if area > mArea {
			mArea = area
		}
		if lower == height[left] {
			left++
		} else {
			right--
		}
	}
	return mArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))  // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                       // Output: 1
	fmt.Println(maxArea([]int{1, 1, 1, 100, 101, 1, 1, 1})) // Output: 100
}
