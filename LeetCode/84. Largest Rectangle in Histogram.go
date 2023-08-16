package main

import "fmt"

type Stack []int

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	var h, p Stack
	maxArea := 0
	for i := 0; i < n; i++ {
		start := i
		for len(h) > 0 && h.Top() > heights[i] {
			maxArea = max(maxArea, h.Top()*(i-p.Top()))
			start = p.Top()
			h.Pop()
			p.Pop()
		}
		h.Push(heights[i])
		p.Push(start)
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
