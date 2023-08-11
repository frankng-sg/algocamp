package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	sum := 0
	max := 0
	for i := 1; i < n; i++ {
		profit := prices[i] - prices[i-1]
		if sum+profit > profit {
			sum = sum + profit
		} else {
			sum = profit
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{1, 2}))    // Output: 1
	fmt.Println(maxProfit([]int{2, 3, 2})) // Output: 1
}
