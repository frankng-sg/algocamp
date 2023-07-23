package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pair := make(map[int]int)

	l := len(nums)
	for i := 0; i < l; i++ {
		if j, ok := pair[nums[i]]; ok {
			return []int{j, i}
		}
		pair[target-nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
