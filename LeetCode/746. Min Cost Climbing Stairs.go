package main

import "fmt"

func lowerInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	for i := 2; i < n; i++ {
		cost[i] = cost[i] + lowerInt(cost[i-1], cost[i-2])
	}
	return lowerInt(cost[n-1], cost[n-2])
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // Output: 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // Output: 6
}
