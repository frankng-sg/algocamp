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

func change2(amount int, coins []int) int {
	count := make(map[[2]int]int)
	n := len(coins)
	var try func(int, int) int
	try = func(index, total int) int {
		if total == amount {
			return 1
		}
		if index == n || total > amount {
			return 0
		}
		if v, ok := count[[2]int{index, total}]; ok {
			return v
		}
		sum := 0
		for i := index; i < n; i++ {
			sum += try(i, total+coins[i])
		}
		count[[2]int{index, total}] = sum
		return sum
	}
	return try(0, 0)
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) // Output: 4
	fmt.Println(change(3, []int{2}))       // Output: 0
}
