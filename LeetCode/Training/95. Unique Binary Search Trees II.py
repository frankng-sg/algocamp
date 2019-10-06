# PROBLEM:
# Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
# Input: 3
# Output:
# [
#   [1,null,3,2],
#   [3,2,null,1],
#   [3,1,null,null,2],
#   [2,1,3],
#   [1,null,2,null,3]
# ]
# Explanation:
# The above output corresponds to the 5 unique BST's shown below:
#
#    1         3     3      2      1
#     \       /     /      / \      \
#      3     2     1      1   3      2
#     /     /       \                 \
#    2     1         2                 3

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def generateTrees(self, n: int):  # return List[TreeNode]
        # Definition for a binary tree node.

        def generate(min_val: int, max_val: int) -> list:
            if min_val > max_val:
                return [None]
            elif min_val == max_val:
                return [TreeNode(min_val)]

            list_of_trees = []
            for val in range(min_val, max_val + 1):
                for left_root in generate(min_val, val - 1):
                    for right_root in generate(val + 1, max_val):
                        root = TreeNode(val)
                        root.left = left_root
                        root.right = right_root
                        list_of_trees.append(root)
            return list_of_trees

        if n is 0:
            return []
        return generate(1, n)
