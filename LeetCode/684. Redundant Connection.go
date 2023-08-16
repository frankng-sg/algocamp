package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	root := func(node int) int {
		for parent[node] != node {
			node = parent[node]
		}
		return node
	}
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		if parent[u] == 0 && parent[v] == 0 {
			parent[u] = u // root
		}
		if parent[v] == 0 {
			parent[v] = root(u)
		} else if parent[u] == 0 {
			parent[u] = root(v)
		} else if root(u) == root(v) {
			return edge
		} else {
			parent[root(v)] = root(u)
		}
	}
	return []int{}
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}})) // Output: [1 4]
	fmt.Println(findRedundantConnection([][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}})) // Output: [2 5]
}
