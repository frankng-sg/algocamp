# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.occurs = {}

    def search(self, node: TreeNode):
        if not node:
            return

        if node.val in self.occurs:
            self.occurs[node.val] += 1
        else:
            self.occurs[node.val] = 1

        self.search(node.left)
        self.search(node.right)

    def findMode(self, root: TreeNode):
        if not root:
            return []

        self.search(root)

        max_val = max(self.occurs.items(), key=lambda x : x[1])[1]
        mode = []
        for key, val in self.occurs.items():
            if val == max_val:
                mode += [key]

        return mode
