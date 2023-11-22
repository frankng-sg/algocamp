package main

import (
	"fmt"
	"sort"
)

// Question
// Given an array of positive integers. Each integer has at least two digits. Find the maximum sum of 2 numbers where the
// first and last digits of first number are same as the first and last digits of second number
//
// Sample Input: [10, 123, 190, 10093]
// Sample Output: 10216
// Explanation:
// There are two possible sums:
// 10 + 190 = 200
// 123 + 10093 = 10216
// Hence, the maximum sum is 10216

func hash(v int) int {
	result := v % 10
	for v >= 10 {
		v /= 10
	}
	result = v*10 + result
	return result
}

func maxSum(arr []int) int {
	sort.Ints(arr)
	m := make(map[int][]int, len(arr)>>1)
	for i := len(arr) - 1; i >= 0; i-- {
		key := hash(arr[i])
		m[key] = append(m[key], arr[i])
	}
	sum := 0
	for _, v := range m {
		if len(v) > 1 && sum < v[0]+v[1] {
			sum = v[0] + v[1]
		}
	}
	return sum
}

func main() {
	fmt.Println(maxSum([]int{10, 123, 190, 10093}))
}
