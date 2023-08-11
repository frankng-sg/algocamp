package main

import "fmt"

func maxProfit(prices []int) int {
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

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // Output: 5
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))    // Output: 4
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // Output: 0
}
