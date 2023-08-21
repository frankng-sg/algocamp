package main

import "math"

func mergeArray(merge, arr1, arr2 []int) {
	n1, n2 := len(arr1), len(arr2)
	i, j := 0, 0
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			merge[i+j] = arr1[i]
			i++
		} else {
			merge[i+j] = arr2[j]
			j++
		}
	}
	for i < n1 {
		merge[i+j] = arr1[i]
		i++
	}
	for j < n2 {
		merge[i+j] = arr2[j]
		j++
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))
	mergeArray(merged, nums1, nums2)
	n := len(merged)
	if n%2 == 0 {
		return float64(merged[n>>1-1]+merged[n>>1]) / 2
	}
	return float64(merged[n/2])
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	imin, imax, mid := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) >> 1
		j := mid - i
		if i < m && nums1[i] < nums2[j-1] {
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			imax = i - 1
		} else {
			maxOfLeft, minOfRight := 0., 0.
			if i == 0 {
				maxOfLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxOfLeft = float64(nums1[i-1])
			} else {
				maxOfLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
				if (m+n)%2 == 1 {
					return maxOfLeft
				}
			}
			if (m+n)%2 == 1 {
				return maxOfLeft
			}
			if i == m {
				minOfRight = float64(nums2[j])
			} else if j == n {
				minOfRight = float64(nums1[i])
			} else {
				minOfRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (maxOfLeft + minOfRight) / 2
		}
	}

	return 0.0
}
