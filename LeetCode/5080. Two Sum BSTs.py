# 5080. Two Sum BSTs
# Given two binary search trees, return True if and only if there is a node in the first tree and a node in
# the second tree whose values sum up to a given integer target.
#
# Input: root1 = [2,1,4], root2 = [1,0,3], target = 5
# Output: true
# Explanation: 2 and 3 sum up to 5.
#
# Input: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
# Output: false

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def binary_search(root, target):
    if not root:
        return False

    if root.val == target:
        return True

    if root.val > target:
        return binary_search(root.left, target)

    return binary_search(root.right, target)


class Solution:
    def twoSumBSTs(self, root1: TreeNode, root2: TreeNode, target: int) -> bool:
        if not root1 and not root2:
            return False
        if not root1:
            return root2.val == target
        if not root2:
            return root1.val == target

        stack = [root1]
        while stack:
            node = stack.pop()

            if binary_search(root2, target - node.val):
                return True
            if node.left:
                stack.append(node.left)
            if node.right:
                stack.append(node.right)

        return False
