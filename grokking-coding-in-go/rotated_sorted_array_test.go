package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func binarySearchRotated(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < target && target <= nums[r] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func Test_binarySearchRotated(t *testing.T) {
	assert.Equal(t, 4, binarySearchRotated([]int{6, 7, 1, 2, 3, 4, 5}, 3))
	assert.Equal(t, 1, binarySearchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	assert.Equal(t, 0, binarySearchRotated([]int{5, 6, 7, 0, 1, 2, 3, 4}, 5))
	assert.Equal(t, -1, binarySearchRotated([]int{5, 6, 7, 0, 1, 2, 3, 4}, -1))
}
