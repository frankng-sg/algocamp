package main

import "fmt"

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])

	var visit func(int, int, *bool)
	visit = func(i, j int, sub *bool) {
		if i < 0 || i >= m || j < 0 || j >= n || grid2[i][j] == 0 {
			return
		}
		if grid1[i][j] == 0 {
			*sub = false
		}
		grid2[i][j] = 0
		visit(i-1, j, sub)
		visit(i, j-1, sub)
		visit(i+1, j, sub)
		visit(i, j+1, sub)
	}
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				sub := true
				if visit(i, j, &sub); sub {
					sum += 1
				}
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(countSubIslands([][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
		[][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}))
}
