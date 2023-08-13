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
	sum := make([]int, n+1)
	sum[0] = 0
	sum[1] = 0
	for i := 2; i <= n; i++ {
		sum[i] = lowerInt(sum[i-1]+cost[i-1], sum[i-2]+cost[i-2])
	}
	return sum[n]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // Output: 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // Output: 6
}
