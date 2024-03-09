package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func containerWithMostWater(height []int) int {
	maxArea := 0
	i, j := 0, len(height)-1
	for i < j {
		w := j - i
		var h int
		if height[i] > height[j] {
			h, j = height[j], j-1
		} else {
			h, i = height[i], i+1
		}
		area := h * w
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func TestWaterContainer(t *testing.T) {
	type TestCase struct {
		In     []int
		Expect int
	}
	tests := []TestCase{
		{In: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Expect: 49},
		{In: []int{1, 5, 1, 1, 9, 1, 9, 1, 5}, Expect: 35},
	}
	for _, tt := range tests {
		actual := containerWithMostWater(tt.In)
		assert.Equal(t, tt.Expect, actual, fmt.Sprintf("input: %v expect: %d actual: %d", tt.In, tt.Expect, actual))
	}
}
