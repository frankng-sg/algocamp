package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	slen := len(s)
	i := 0
	for i < slen && s[i] == ' ' {
		i++
	}
	if i >= slen {
		return 0
	}
	sign := 1
	limit := math.MaxInt32
	if s[i] == '-' {
		i++
		sign = -1
		limit = -math.MinInt32
	} else if s[i] == '+' {
		i++
	}

	num := 0
	for i < slen {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		digit := int(s[i] - '0')
		if num > (limit-digit)/10 {
			return int(sign * limit)
		}
		num = num*10 + digit
		i++
	}
	return sign * num
}

func main() {
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("+-42"))
	fmt.Println(myAtoi("-4.2"))
	fmt.Println(myAtoi(""))
}
