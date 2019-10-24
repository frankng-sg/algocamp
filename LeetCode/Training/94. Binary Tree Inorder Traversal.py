# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def inorderTraversal(self, root):
        sequence = []

        def inorder(root):
            if root is None:
                return
            inorder(root.left)
            sequence.append(root.val)
            inorder(root.right)

        inorder(root)
        return sequence
