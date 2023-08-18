package main

import "fmt"

func change(amount int, coins []int) int {
	count := make([]int, amount+1)
	count[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				count[j] += count[j-coins[i]]
			}

		}
	}
	return count[amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) // Output: 4
	fmt.Println(change(3, []int{2}))       // Output: 0
}
