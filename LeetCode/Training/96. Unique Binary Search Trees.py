# PROBLEM:
# Given an integer n, count number of unique BST's (binary search trees) that store values 1 ... n.
# Input: 3
# Output: 5


class Solution:
    def numTrees(self, n: int) -> int:
        def count_bst(min_val, max_val):
            if min_val >= max_val:
                return 1
            no_of_trees = 0
            for i in range(min_val, max_val + 1):
                no_of_trees += count_bst(min_val, i - 1) * count_bst(i + 1, max_val)
            return no_of_trees

        return count_bst(1, n)
