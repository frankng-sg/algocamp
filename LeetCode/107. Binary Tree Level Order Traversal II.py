from collections import deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrderBottom(self, root: TreeNode):
        output = []
        if not root:
            return output
        queue = deque()
        queue.append(root)
        while queue:
            level_size = len(queue)
            level_nodes = []
            for i in range(level_size):
                node = queue.popleft()
                level_nodes.append(node.val)
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            if not level_nodes:
                break
            output.append(level_nodes)

        output.reverse()
        return output
