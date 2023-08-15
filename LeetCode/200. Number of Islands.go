package main

import "fmt"

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	sum := 0
	var visit func(int, int)
	visit = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		visit(i+1, j)
		visit(i-1, j)
		visit(i, j-1)
		visit(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				visit(i, j)
				sum += 1
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})) // Output: 1
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})) // Output: 3
}
