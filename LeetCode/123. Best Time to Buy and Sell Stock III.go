package main

import "fmt"

func bestTrade(prices []int) int {
	n := len(prices)
	entry := prices[0]
	max := 0
	for i := 1; i < n; i++ {
		if prices[i] < entry {
			entry = prices[i]
		} else if prices[i]-entry > max {
			max = prices[i] - entry
		}
	}
	return max
}

func maxProfit(prices []int) int {
	n := len(prices)
	maxProfit := make([]int, n)
	maxProfit[0] = 0
	entry := prices[0]
	max := 0
	for i := 1; i < n; i++ {
		if prices[i] < entry {
			entry = prices[i]
		} else if prices[i]-entry > max {
			max = prices[i] - entry
		}
		maxProfit[i] = max
	}
	maxPrice := make([]int, n+1)
	maxPrice[n] = 0
	max = 0
	for i := n - 1; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		maxPrice[i] = max
	}
	max = 0
	for i := 1; i < n; i++ {
		twoTrades := maxProfit[i] + maxPrice[i] - prices[i]
		if twoTrades > max {
			max = twoTrades
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})) // Output: 6
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 5, 4, 6}))    // Output: 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))          // Output: 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))          // Output: 0
}
