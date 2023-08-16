package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	next := make([]int, n)
	next[n-1] = -1
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for next[j] != -1 && temperatures[i] >= temperatures[j] {
			j = next[j]
		}
		if temperatures[i] < temperatures[j] {
			next[i] = j
		} else {
			next[i] = -1
		}
	}
	for i := 0; i < n; i++ {
		if next[i] == -1 {
			next[i] = 0
		} else {
			next[i] = next[i] - i
		}
	}
	return next
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // Output: [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // Output: [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30, 10, 20, 30}))                 // Output: [0,1,1,0]
}
