# Given a binary tree, determine if it is a valid binary search tree (BST).
#
# Assume a BST is defined as follows:
#
# The left subtree of a node contains only nodes with keys less than the node's key.
# The right subtree of a node contains only nodes with keys greater than the node's key.
# Both the left and right subtrees must also be binary search trees.

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def isValidBST(self, root: TreeNode, min_val=None, max_val=None) -> bool:

        if root is None:
            return True
        if min_val is not None and root.val <= min_val:
            return False
        if max_val is not None and root.val >= max_val:
            return False

        if self.isValidBST(root.left, min_val, root.val) is False:
            return False
        if self.isValidBST(root.right, root.val, max_val) is False:
            return False

        return True
