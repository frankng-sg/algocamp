package main

import "fmt"

func productExceptSelf(nums []int) []int {
	nlen := len(nums)
	m := make([]int, nlen)
	mult := 1
	for i := nlen - 1; i >= 0; i-- {
		m[i] = mult
		mult *= nums[i]
	}
	mult = nums[0]
	for i := 1; i < nlen; i++ {
		m[i] = mult * m[i]
		mult *= nums[i]
	}
	return m
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))      // Output: [24,12,8,6]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3})) // Output: [0,0,9,0,0]
}
