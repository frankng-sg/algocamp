package main

import "fmt"

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	var sum [300][5001]int
	n := len(coins)
	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j%coins[i] == 0 {
				sum[i][j] = 1
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= amount; j++ {
			sum[i][j] += sum[i-1][j]
			k := j
			for k-coins[i] > 0 {
				k -= coins[i]
				sum[i][j] += sum[i-1][k]
			}
		}
	}
	return sum[n-1][amount]
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) // Output: 4
	fmt.Println(change(3, []int{2}))       // Output: 0
}
