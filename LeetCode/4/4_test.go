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
		{[]int{1, 3}, []int{2}, 2.00000},
		{[]int{1, 2}, []int{3, 4}, 2.50000},
		{[]int{0, 0}, []int{0, 0}, 0.00000},
		{[]int{}, []int{1}, 1.00000},
		{[]int{2}, []int{}, 2.00000},
		{[]int{1, 3, 5}, []int{2, 4, 6}, 3.50000},
	}

	for _, test := range tests {
		got := findMedianSortedArrays(test.nums1, test.nums2)
		if math.Abs(got-test.want) > 0.00001 {
			t.Errorf("findMedianSortedArrays(%v, %v) = %f; want %f", test.nums1, test.nums2, got, test.want)
		}
	}
}
