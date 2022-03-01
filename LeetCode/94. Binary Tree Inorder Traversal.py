# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def inorderTraversal(self, root):
        if root is None:
            return []

        tree = [root]
        sequence = []
        while tree:
            if tree[-1].left:
                node = tree[-1].left
                tree[-1].left = None
                tree.append(node)
            else:
                node = tree[-1]
                sequence.append(node.val)
                tree.pop()
                if node.right:
                    tree.append(node.right)

        return sequence
