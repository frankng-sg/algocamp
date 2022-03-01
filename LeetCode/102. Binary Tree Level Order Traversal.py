# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def levelOrder(self, root: TreeNode):
        if root is None:
            return []
        queue = [root]
        output = []
        while True:
            same_level_nodes = len(queue)
            list_of_nodes = []
            while same_level_nodes > 0:
                node = queue.pop(0)
                list_of_nodes.append(node.val)
                same_level_nodes -= 1
                if node.left is not None:
                    queue.append(node.left)
                if node.right is not None:
                    queue.append(node.right)
            if not list_of_nodes:
                break

            output.append(list_of_nodes)

        return output
