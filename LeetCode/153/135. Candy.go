package main

import "fmt"

func candy(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			c[i] = max(c[i], c[i-1]+1)
		}
	}
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			c[i-1] = max(c[i-1], c[i]+1)
		}
	}
	sum := 0
	for _, v := range c {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))                // 5
	fmt.Println(candy([]int{1, 2, 2}))                // 4
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1})) // 13
}
