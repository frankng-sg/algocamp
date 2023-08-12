package main

import "fmt"

func lower(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	n := len(height)
	mLeft := 0
	mArea := 0
	for i := 0; i < n; i++ {
		if height[i] > mLeft {
			mLeft = height[i]
			mRight := 0
			for j := n - 1; j > i; j-- {
				if height[j] > mRight {
					mRight = height[j]
					area := lower(mLeft, mRight) * (j - i)
					if area > mArea {
						mArea = area
					}
				}
			}
		}
	}
	return mArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // Output: 49
	fmt.Println(maxArea([]int{1, 1}))                      // Output: 1
}
