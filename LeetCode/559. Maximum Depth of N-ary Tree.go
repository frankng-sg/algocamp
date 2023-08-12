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
		depth := maxDepth(v)
		if max < depth {
			max = depth
		}
	}
	return max + 1
}
