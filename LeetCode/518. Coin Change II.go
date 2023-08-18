package main

import "fmt"

func change(amount int, coins []int) int {
	count := make([]int, amount+1)
	count[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			count[i] += count[i-coin]
		}
	}
	return count[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) // Output: 4
	fmt.Println(change(3, []int{2}))       // Output: 0
}
