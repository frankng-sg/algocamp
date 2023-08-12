package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	nlen := len(nums)
	if nlen <= 0 {
		return nil
	}
	m := 0
	for i := 1; i < nlen; i++ {
		if nums[m] < nums[i] {
			m = i
		}
	}
	return &TreeNode{
		Val:   nums[m],
		Left:  constructMaximumBinaryTree(nums[:m]),
		Right: constructMaximumBinaryTree(nums[m+1:]),
	}
}

func main() {
	fmt.Println(constructMaximumBinaryTree([]int{1, 2, 3}))
}
