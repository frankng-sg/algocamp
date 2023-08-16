package main

import "fmt"

func largestRectangleArea(heights []int) int {
	maxh := []int{0}
	pos := []int{0}
	n := len(heights)
	maxArea := 0
	for i := 0; i < n; i++ {
		if heights[i] > maxArea {
			maxArea = heights[i]
		}
		found := false
		for j := len(maxh) - 1; j >= 0; j-- {
			if heights[i] < maxh[j] {
				found = true
				maxh[j] = heights[i]
			}
			area := maxh[j] * (i - pos[j] + 1)
			if area > maxArea {
				maxArea = area
			}
		}
		if !found {
			maxh = append(maxh, heights[i])
			pos = append(pos, i)
		}
	}
	return maxArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{1}))                               // Output: 1
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))                // Output: 10
	fmt.Println(largestRectangleArea([]int{2, 4}))                            // Output: 4
	fmt.Println(largestRectangleArea([]int{2, 1, 1, 1, 1, 5, 5, 1, 1, 1, 1})) // Output: 11
}
