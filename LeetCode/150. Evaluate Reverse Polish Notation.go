package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	num := []int{}
	i := 0
	for _, v := range tokens {
		switch v {
		case "+":
			i--
			num[i-1] += num[i]
		case "-":
			i--
			num[i-1] -= num[i]
		case "*":
			i--
			num[i-1] *= num[i]
		case "/":
			i--
			num[i-1] /= num[i]
		default:
			k, _ := strconv.Atoi(v)
			if i == len(num) {
				num = append(num, k)
			} else {
				num[i] = k
			}
			i++
		}
	}
	return num[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))                                             // Output: 9
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))                                            // Output: 6
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})) // Output: 2
}
