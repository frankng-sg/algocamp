package main

import "fmt"

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func higherInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	mLeft := make([]int, n)
	mRight := make([]int, n)
	for i := 1; i < n; i++ {
		mLeft[i] = higherInt(mLeft[i-1], height[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		mRight[i] = higherInt(mRight[i+1], height[i+1])
	}
	sum := 0
	for i := 1; i < n-1; i++ {
		if height[i] < mLeft[i] && height[i] < mRight[i] {
			sum += lowerInt(mLeft[i], mRight[i]) - height[i]
		}
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // Output: 6
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 1, 2, 1}))    // Output: 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // Output: 9
	fmt.Println(trap([]int{4, 2, 3}))                            // Output: 1
}
