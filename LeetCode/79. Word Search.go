package main

import "fmt"

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])
	lword := len(word)

	// Initialize the visited 2D slice with false values
	visited := make([][]bool, height)
	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			visited[i][j] = false
		}
	}

	var find func(i, j, d int) bool
	find = func(i, j, d int) bool {
		if d == lword {
			return true
		}
		if i < 0 || i >= height || j < 0 || j >= width || visited[i][j] {
			return false
		}
		if board[i][j] != word[d] {
			return false
		}
		visited[i][j] = true
		if find(i+1, j, d+1) || find(i-1, j, d+1) || find(i, j+1, d+1) || find(i, j-1, d+1) {
			return true
		}
		visited[i][j] = false
		return false
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if find(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Printf("%v\n", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB")) // false
	fmt.Printf("%v\n", exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABC"))  // true
}
