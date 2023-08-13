package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	result := make([]int, n+1)
	c := 1
	for i := n - 1; i >= 0; i-- {
		sum := c + digits[i]
		if sum >= 10 {
			result[i+1] = sum - 10
			c = 1
		} else {
			result[i+1] = sum
			c = 0
		}
	}
	result[0] = c
	if c > 0 {
		return result
	} else {
		return result[1:]
	}
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9})) // Output: [1, 0, 0, 0]
}
