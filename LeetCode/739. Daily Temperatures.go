package main

import (
	"fmt"
)

func dailyTemperatures2(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}
	top := -1
	for i, t := range temperatures {
		for top >= 0 && temperatures[stack[top]] < t {
			loc := stack[top]
			ans[loc] = i - loc
			stack = stack[:top]
			top--
		}
		stack = append(stack, i)
		top++
	}
	return ans
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	count := make([]int, n)
	count[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for count[j] != 0 && temperatures[i] >= temperatures[j] {
			j += count[j]
		}
		if temperatures[i] < temperatures[j] {
			count[i] = j - i
		}
	}
	return count
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})) // Output: [1,1,4,2,1,1,0,0]
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))                 // Output: [1,1,1,0]
	fmt.Println(dailyTemperatures([]int{30, 10, 20, 30}))                 // Output: [0,1,1,0]
}
