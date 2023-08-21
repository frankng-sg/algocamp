package main

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{}, []int{1}, 1.},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},
		{[]int{1, 2, 7}, []int{3, 4, 6}, 3.5},
		{[]int{1, 3}, []int{2}, 2.},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.},
		{[]int{2}, []int{}, 2.},
		{[]int{1, 3, 5}, []int{2, 4, 6}, 3.5},
		{[]int{1, 2, 6}, []int{7, 8, 9}, 6.5},
		{[]int{1, 12}, []int{6, 7}, 6.5},
	}

	for _, test := range tests {
		got := findMedianSortedArrays(test.nums1, test.nums2)
		if math.Abs(got-test.want) > 0.00001 {
			t.Errorf("findMedianSortedArrays(%v, %v) = %f; want %f", test.nums1, test.nums2, got, test.want)
		}
	}
}
