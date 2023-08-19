package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	rev := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if rev > (math.MaxInt32-digit)/10 || rev < (math.MinInt32-digit)/10 {
			return 0
		}
		rev = rev*10 + digit
	}
	return rev
}

func main() {
	fmt.Println(reverse(1123123129))  // Output: 0
	fmt.Println(reverse(1123123121))  // Output: 1213213211
	fmt.Println(reverse(-1123123129)) // Output: 0
	fmt.Println(reverse(-1123123121)) // Output: -1213213211
}
