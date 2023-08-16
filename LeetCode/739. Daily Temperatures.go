package main

import (
	"fmt"
	"sort"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	pos := make(map[int][]int)
	for i, v := range temperatures {
		pos[v] = append(pos[v], i)
	}
	for t := 30; t <= 100; t++ {
		sort.Ints(pos[t])
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		min := n
		for t := temperatures[i] + 1; t <= 100; t++ {
			for _, loc := range pos[t] {
				if loc > i && loc-i < min {
					min = loc - i
				}
			}

		}
		if min < n {
			ans[i] = min
		}
	}
	return ans
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // Output: [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // Output: [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30, 10, 20, 30}))                 // Output: [0,1,1,0]
}
