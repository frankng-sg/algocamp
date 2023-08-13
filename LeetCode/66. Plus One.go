package main

import "fmt"

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 && digits[i] == 9 {
		digits[i] = 0
		i--
	}
	if i >= 0 {
		digits[i]++
	} else {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9})) // Output: [1, 0, 0, 0]
}
