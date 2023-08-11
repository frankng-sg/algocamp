package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	open := prices[0]
	close := prices[0]
	sum := 0
	for i := 1; i < n; i++ {
		if prices[i] < close {
			sum += close - open
			open = prices[i]
			close = prices[i]
		} else {
			close = prices[i]
		}
	}
	return sum + close - open
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // Output: 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    // Output: 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // Output: 0
}
