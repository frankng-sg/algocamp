package main

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
