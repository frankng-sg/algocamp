package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := []int{1}
	l := 0
	r := 0
	for l <= r {
		node := q[l]
		if node.Left == nil && node.Right == nil {
			return depth[l]
		}
		if node.Left != nil {
			q = append(q, node.Left)
			r++
			depth = append(depth, depth[l]+1)
		}
		if node.Right != nil {
			q = append(q, node.Right)
			r++
			depth = append(depth, depth[l]+1)
		}
		l++
	}
	return depth[r]
}
