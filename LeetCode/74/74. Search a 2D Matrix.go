package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	nrow, ncol := len(matrix), len(matrix[0])
	l, r := 0, nrow*ncol-1
	for l <= r {
		mid := l + (r-l)>>1
		row, col := mid/ncol, mid%ncol
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}}, 23)) // Output: false
	fmt.Println(searchMatrix([][]int{
		{1, 11, 16, 20},
		{21, 23, 24, 25},
		{53, 57, 58, 99},
	}, 23)) // Output: true
	fmt.Println(searchMatrix([][]int{
		{1, 11, 16, 20},
		{21, 23, 24, 25},
		{53, 57, 58, 99},
	}, 22)) // Output: false
	fmt.Println(searchMatrix([][]int{
		{-8, -6, -6, -4, -2, -1, -1, -1, 0, 1, 2, 4, 6, 7, 7},
		{10, 10, 12, 13, 13, 14, 14, 16, 17, 17, 17, 17, 17, 18, 19},
		{22, 24, 26, 28, 29, 31, 32, 34, 34, 34, 34, 36, 38, 39, 39},
		{40, 40, 41, 43, 43, 44, 46, 47, 49, 49, 50, 52, 53, 55, 55},
		{56, 57, 59, 61, 62, 64, 65, 65, 66, 67, 68, 68, 69, 70, 71},
		{74, 75, 77, 77, 79, 79, 79, 80, 81, 83, 85, 87, 88, 89, 89},
		{91, 93, 94, 96, 97, 98, 99, 99, 100, 100, 102, 104, 105, 107, 107},
		{110, 112, 114, 114, 115, 117, 117, 118, 118, 120, 120, 121, 123, 124, 124},
		{127, 127, 129, 131, 133, 134, 134, 135, 137, 139, 139, 140, 142, 143, 144},
		{145, 146, 147, 149, 151, 151, 153, 155, 156, 156, 158, 160, 162, 162, 163},
	}, 107)) // Output: true
}
