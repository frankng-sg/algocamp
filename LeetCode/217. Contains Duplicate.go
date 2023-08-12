package main

import "fmt"

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool)
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			return true
		}
		exist[v] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // Output: true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // Output: false
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // Output: true
}
