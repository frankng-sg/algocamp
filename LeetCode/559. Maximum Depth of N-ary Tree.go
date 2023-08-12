package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, v := range root.Children {
		if depth := maxDepth(v); max < depth {
			max = depth
		}
	}
	return max + 1
}
