package main

import (
	"testing"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		piles []int
		h     int
		want  int
	}{
		{[]int{312884470}, 968709470, 1},
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{1, 1, 1, 6, 7, 11}, 7, 7},
		{[]int{1, 1, 1, 6, 7, 11}, 6, 11},
	}

	for _, test := range tests {
		got := minEatingSpeed(test.piles, test.h)
		if got != test.want {
			t.Errorf("minEatingSpeed(%v, %v) = %v; want %v", test.piles, test.h, got, test.want)
		}
	}
}
