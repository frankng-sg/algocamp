package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfitUnlimitedTrades(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

// Time: O(k * n), Space: O(n)
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > n>>1 {
		return maxProfitUnlimitedTrades(prices)
	}
	prev := make([]int, n)
	curr := make([]int, n)
	for i := 1; i <= k; i++ {
		maxDiff := -prices[0] // buy at prices[0]
		curr[0] = 0
		for j := 1; j < n; j++ {
			curr[j] = max(curr[j-1], prices[j]+maxDiff) // sell at prices[j]
			maxDiff = max(maxDiff, prev[j]-prices[j])   // buy at prices[j]
		}
		copy(prev, curr)
	}

	return curr[n-1]
}

func main() {
	fmt.Println(maxProfit(6, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0})) // Output: 15
	fmt.Println(maxProfit(2, []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0})) // Output: 13
	fmt.Println(maxProfit(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))       // Output: 6
	fmt.Println(maxProfit(2, []int{2, 4, 1}))                      // Output: 2
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))             // Output: 7
	fmt.Println(maxProfit(2, []int{1}))                            // Output: 0
}
