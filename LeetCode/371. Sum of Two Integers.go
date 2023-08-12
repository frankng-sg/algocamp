package main

import "fmt"

// Given two integers a and b, return the sum
// of the two integers without using the operators + and -.
func getSum(a, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	fmt.Println(getSum(3, 5))   // Output: 8
	fmt.Println(getSum(-3, 5))  // Output: 2
	fmt.Println(getSum(-3, -5)) // Output: -8
}
