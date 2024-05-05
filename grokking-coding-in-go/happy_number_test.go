package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func next(num int) int {
	sum := 0
	for num > 0 {
		sum += (num % 10) * (num % 10)
		num /= 10
	}
	return sum
}

func isHappy(num int) bool {
	fast, slow := next(num), num
	for fast != 1 && fast != slow {
		fast = next(next(fast))
		slow = next(slow)
	}
	return fast == 1
}

func Test_isHappy(t *testing.T) {
	assert.True(t, isHappy(19))
	assert.False(t, isHappy(2999999999))
}
