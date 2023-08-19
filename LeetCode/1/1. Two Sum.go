package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pair := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := pair[v]; ok {
			return []int{j, i}
		}
		pair[target-v] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
