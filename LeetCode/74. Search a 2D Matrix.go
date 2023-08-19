package main

import "fmt"

func bsearch(n, target int, getVal func(int) int) int {
	l, r := 0, n
	for l < r-1 {
		mid := l + (r-l)>>1
		if getVal(mid) <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func searchMatrix(matrix [][]int, target int) bool {
	nrow := len(matrix)
	ncol := len(matrix[0])

	getFirstValueInRow := func(i int) int {
		return matrix[i][0]
	}

	// find which row the target should be in
	row := bsearch(nrow, target, getFirstValueInRow)
	getValueInRow := func(i int) int {
		return matrix[row][i]
	}
	col := bsearch(ncol, target, getValueInRow)

	return matrix[row][col] == target
}

func main() {
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
}
