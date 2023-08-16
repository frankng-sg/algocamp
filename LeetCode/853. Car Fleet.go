package main

import (
	"fmt"
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	loc := make([]int, n)
	for i := 0; i < n; i++ {
		loc[i] = i
	}
	sort.Slice(loc, func(i, j int) bool {
		return position[loc[i]] < position[loc[j]]
	})
	fleets := 1
	tfleet := float64(target-position[loc[n-1]]) / float64(speed[loc[n-1]])
	for i := n - 2; i >= 0; i-- {
		t := float64(target-position[loc[i]]) / float64(speed[loc[i]])
		if t > tfleet {
			fleets++
			tfleet = t
		}
	}
	return fleets
}

func main() {
	fmt.Println(carFleet(17, []int{8, 12, 16, 11, 7}, []int{6, 9, 10, 9, 7})) // Output: 4
	fmt.Println(carFleet(12, []int{0, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))     // Output: 3
	fmt.Println(carFleet(10, []int{0, 4, 2}, []int{2, 1, 3}))                 // Output: 1
}
