package main

import "fmt"

// Time: O(n^2) Space: O(1)
func rotate(matrix [][]int) {
	n := len(matrix)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		for i := l; i < r; i++ {
			tmp := matrix[l][i]
			matrix[l][i] = matrix[r-i+l][l]
			matrix[r-i+l][l] = matrix[r][r-i+l]
			matrix[r][r-i+l] = matrix[i][r]
			matrix[i][r] = tmp
		}
	}
}

func main() {
	var m [][]int

	m = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(m)
	fmt.Println(m) // Output: [[7,4,1],[8,5,2],[9,6,3]]

	m = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m)
	fmt.Println(m) // Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

}
