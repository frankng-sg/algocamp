class Solution:
    def __init__(self):
        self.used_digit_col = [set() for i in range(9)]
        self.used_digit_row = [set() for i in range(9)]
        self.used_digit_zone = [set() for i in range(9)]

    def solveSudoku(self, board) -> None:
        """
        Do not return anything, modify board in-place instead.
        """

        def backtrack(board, row, col):
            if row >= len(board):
                return True

            # iterate all possible candidates.
            if col >= len(board[0]) - 1:
                next_col = 0
                next_row = row + 1
            else:
                next_col = col + 1
                next_row = row

            if board[row][col] != '.':
                return backtrack(board, next_row, next_col)

            for candidate in '123456789':
                if candidate not in self.used_digit_col[col] \
                        and candidate not in self.used_digit_row[row] \
                        and candidate not in self.used_digit_zone[(row // 3) * 3 + (col // 3)]:
                    # try this partial candidate solution
                    board[row][col] = candidate
                    self.used_digit_col[col].add(candidate)
                    self.used_digit_row[row].add(candidate)
                    self.used_digit_zone[(row // 3) * 3 + (col // 3)].add(candidate)

                    # given the candidate, explore further.
                    if backtrack(board, next_row, next_col):
                        return True

                    # backtrack
                    board[row][col] = '.'
                    self.used_digit_col[col].remove(candidate)
                    self.used_digit_row[row].remove(candidate)
                    self.used_digit_zone[(row // 3) * 3 + (col // 3)].remove(candidate)
            return False

        for i in range(0, len(board)):
            for j in range(0, len(board[0])):
                self.used_digit_col[j].add(board[i][j])
                self.used_digit_row[i].add(board[i][j])
                self.used_digit_zone[(i // 3) * 3 + (j // 3)].add(board[i][j])

        backtrack(board, 0, 0)
        return board


board = [["5", "3", ".", ".", "7", ".", ".", ".", "."],
         ["6", ".", ".", "1", "9", "5", ".", ".", "."],
         [".", "9", "8", ".", ".", ".", ".", "6", "."],
         ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
         ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
         ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
         [".", "6", ".", ".", ".", ".", "2", "8", "."],
         [".", ".", ".", "4", "1", "9", ".", ".", "5"],
         [".", ".", ".", ".", "8", ".", ".", "7", "9"]]

func = Solution()
func.solveSudoku(board)
for i in range(0, 9):
    print(board[i])
