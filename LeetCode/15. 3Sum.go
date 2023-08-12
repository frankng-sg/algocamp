package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	nlen := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	count := make(map[int]int)
	for _, v := range nums {
		count[v] += 1
	}
	exist := make(map[[3]int]bool)
	for i := 0; i < nlen; i++ {
		count[nums[i]]--
		for j := i + 1; j < nlen; j++ {
			count[nums[j]]--
			v := -nums[i] - nums[j]
			if count[v] > 0 && !exist[[3]int{nums[i], nums[j], v}] {
				result = append(result, []int{nums[i], nums[j], v})
				exist[[3]int{nums[i], nums[j], v}] = true
			}
		}
		for j := i + 1; j < nlen; j++ {
			count[nums[j]]++
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
