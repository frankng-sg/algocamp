package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}

func main() {
	nums := []int{3, 2, 2, 3}
	k := removeElement(nums, 3)
	fmt.Println(k, nums[:k]) // Output: 2 [2, 2]
}
