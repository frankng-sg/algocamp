package main

import "fmt"

func higherInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func trap(height []int) int {
	sum := 0
	n := len(height)
	lMax := height[0]
	rMax := height[n-1]
	for l, r := 0, n-1; l < r; {
		if lMax < rMax {
			l++
			lMax = higherInt(lMax, height[l])
			sum += lMax - height[l]
		} else {
			r--
			rMax = higherInt(rMax, height[r])
			sum += rMax - height[r]
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
