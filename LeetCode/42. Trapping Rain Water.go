package main

import "fmt"

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	n := len(height)
	sum := 0
	for i := 0; i < n-1; {
		m := i + 1
		for j := i + 1; j < n; j++ {
			if height[j] > height[m] {
				m = j
			}
			if height[j] >= height[i] {
				break
			}
		}
		lim := lowerInt(height[i], height[m])
		for j := i + 1; j < m; j++ {
			sum += lim - height[j]
		}
		i = m
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // Output: 6
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 1, 2, 1}))    // Output: 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // Output: 9
	fmt.Println(trap([]int{4, 2, 3}))                            // Output: 1
}
