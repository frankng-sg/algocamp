# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth(self, root):

        def search(stack, max_depth):
            node, depth = stack.pop()

            if (max_depth < depth):
                max_depth = depth

            if (node.left != None):
                stack.append([node.left, depth + 1])
            if (node.right != None):
                stack.append([node.right, depth + 1])

            if (stack == []):
                return max_depth
            else:
                return search(stack, max_depth)

        if (root == None):
            return 0
        stack = [[root, 1]]
        return search(stack, 1)


        
