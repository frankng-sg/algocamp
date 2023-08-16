package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	nums := []int{}
	calc := func(op string) {
		n := len(nums)
		switch op {
		case "+":
			nums[n-2] += nums[n-1]
		case "-":
			nums[n-2] -= nums[n-1]
		case "*":
			nums[n-2] *= nums[n-1]
		case "/":
			nums[n-2] /= nums[n-1]
		}
		nums = nums[:n-1]
	}
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			calc(v)
		} else {
			k, _ := strconv.Atoi(v)
			nums = append(nums, k)
		}
	}
	return nums[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // Output: 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // Output: 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // Output: 2
}
