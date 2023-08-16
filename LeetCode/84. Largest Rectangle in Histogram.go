package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	h, p := []int{0}, []int{0}
	maxArea := 0
	for i := 0; i < n; i++ {
		start := i
		j := len(h) - 1
		for ; j >= 0 && h[j] > heights[i]; j-- {
			maxArea = max(maxArea, h[j]*(i-p[j]))
			start = p[j]
		}
		h = h[:j+1]
		p = p[:j+1]
		if heights[i] > h[j] {
			h = append(h, heights[i])
			p = append(p, start)
		}
	}
	for i := 0; i < len(h); i++ {
		maxArea = max(maxArea, h[i]*(n-p[i]))
	}
	return maxArea
}

func main() {
	fmt.Println(largestRectangleArea([]int{9, 0}))                                                 // Output: 9
	fmt.Println(largestRectangleArea([]int{1}))                                                    // Output: 1
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))                                     // Output: 10
	fmt.Println(largestRectangleArea([]int{2, 4}))                                                 // Output: 4
	fmt.Println(largestRectangleArea([]int{2, 1, 1, 1, 1, 5, 5, 1, 1, 1, 1}))                      // Output: 11
	fmt.Println(largestRectangleArea([]int{2, 1, 1, 1, 1, 5, 5, 9, 3, 2, 5, 7, 4, 4, 3, 6, 7, 8})) // Output: 26
}
