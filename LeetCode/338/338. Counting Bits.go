package main

import "fmt"

func countBits(n int) []int {
	output := make([]int, n+1)
	for i := 1; i <= n; i++ {
		output[i] = output[i>>1] + (i & 1)
	}
	return output
}

func main() {
	fmt.Println(countBits(5)) // Output: [0,1,1,2,1,2]
	fmt.Println(countBits(2)) // Output: [0,1,1]
}
