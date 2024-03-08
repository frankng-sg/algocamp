package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func atoi(str string) int32 {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	sign := int32(1)
	if str[i] == '-' {
		sign, i = -1, i+1
	}
	result := int32(0)
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		digit := int32(str[i] - '0')
		newResult := result*10 + digit
		if newResult < 0 { // overflow
			switch sign {
			case 1:
				return math.MaxInt32
			case -1:
				return math.MinInt32
			}
		}
		result = newResult
		i++
	}
	return result * sign
}

func myAtoi(str string) int {
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}
	sign := 1
	if str[i] == '-' {
		sign, i = -1, i+1
	}
	result := 0
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		digit := int(str[i] - '0')
		newResult := result*10 + digit
		if newResult < 0 { // overflow
			switch sign {
			case 1:
				return math.MaxInt
			case -1:
				return math.MinInt
			}
		}
		result = newResult
		i++
	}
	return result * sign
}

func myAtoiSol(s string) int {
	result := 0
	sign := 1
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + digit
		i++
	}

	return sign * result
}

func TestAtoi(t *testing.T) {
	type TestCase struct {
		In     string
		Expect int
	}
	tests := []TestCase{
		//{In: "25 you rock", Expect: 25},
		//{In: "  -54.4  ", Expect: -54},
		//{In: "87.5", Expect: 87},
		//{In: "-00035", Expect: -35},
		//{In: "-minus86", Expect: 0},
		{In: "-2147483648", Expect: -2147483648},
		{In: "2147483648", Expect: 2147483647},
	}
	for _, tt := range tests {
		actual := myAtoiSol(tt.In)
		assert.Equal(t, tt.Expect, actual, fmt.Sprintf("input: %s expect: %d actual: %d", tt.In, tt.Expect, actual))
	}
}
