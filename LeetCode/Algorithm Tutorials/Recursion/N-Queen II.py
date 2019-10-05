# N-Queens II
# The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
# Given an integer n, return the number of distinct solutions to the n-queens puzzle.
#
# Example:
#
# Input: 4
# Output: 2
# Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
#
#  [".Q..",  // Solution 1
#   "...Q",
#   "Q...",
#   "..Q."],
#
#  ["..Q.",  // Solution 2
#   "Q...",
#   "...Q",
#   ".Q.."]
#
MAX_N = 100


class Solution:
    def __init__(self):
        self.occupied_row = [False] * 2 * MAX_N
        self.occupied_left_diagonal = [False] * 2 * MAX_N
        self.occupied_right_diagonal = [False] * 2 * MAX_N
        self.no_of_solutions = 0

    def recur_search(self, size: int, col: int, occupied_row, occupied_xy_sum, occupied_xy_sub):
        if col < 0:
            self.no_of_solutions += 1
            return

        for row in range(size):
            if row not in occupied_row and row + col not in occupied_xy_sum and row - col not in occupied_xy_sub:
                self.recur_search(size, col - 1, occupied_row + [row], occupied_xy_sum + [row + col],
                                  occupied_xy_sub + [row - col])

    def totalNQueens(self, n: int) -> int:
        self.recur_search(n, n - 1, [], [], [])
        return self.no_of_solutions


func = Solution()
print(func.totalNQueens(4))
