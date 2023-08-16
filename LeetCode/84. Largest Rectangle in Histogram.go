package main

import "fmt"

func higherInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	maxh := []int{}
	pos := []int{}
	n := len(heights)
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := len(maxh) - 1; j >= 0; j-- {
			if heights[i] < maxh[j] {
				maxh[j] = heights[i]
			}
		}
		for j := len(maxh) - 2; j >= 0; j-- {
			if maxh[j] >= maxh[j+1] {
				maxh = maxh[:j+1]
				pos = pos[:j+1]
			}
		}
		if len(maxh) == 0 || heights[i] > maxh[len(maxh)-1] {
			maxh = append(maxh, heights[i])
			pos = append(pos, i)
		}
		for j := len(maxh) - 1; j >= 0; j-- {
			maxArea = higherInt(maxArea, maxh[j]*(i-pos[j]+1))
		}
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
