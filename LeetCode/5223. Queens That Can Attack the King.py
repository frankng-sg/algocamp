class Solution:
    def queensAttacktheKing(self, queens, king):
        board = []
        for i in range(8):
            board.append([-1, -1, -1, -1, -1, -1, -1, -1])
        for i in range(len(queens)):
            queen = queens[i]
            board[queen[0]][queen[1]] = i
        Qlist = []
        # Check Column
        row = king[0]
        while row > 0:
            row -= 1
            Qindex = board[row][king[1]]
            if Qindex != -1:
                Qlist.append(queens[Qindex])
                break

        row = king[0]
        while row < 7:
            row += 1
            Qindex = board[row][king[1]]
            if Qindex != -1:
                Qlist.append(queens[Qindex])
                break

        # Check Row
        col = king[1]
        while col > 0:
            col -= 1
            Qindex = board[king[0]][col]
            if Qindex != -1:
                Qlist.append(queens[Qindex])
                break

        col = king[1]
        while col < 7:
            col += 1
            Qindex = board[king[0]][col]
            if Qindex != -1:
                Qlist.append(queens[Qindex])
                break

        # Check Diagonal
        # Upper Right
        row = king[0]
        col = king[1]
        while row > 0 and col < 7:
            row -= 1
            col += 1
            Qindex = board[row][col]
            if board[row][col] != -1:
                Qlist.append(queens[Qindex])
                break

        # Upper Left
        row = king[0]
        col = king[1]
        while row > 0 and col > 0:
            row -= 1
            col -= 1
            Qindex = board[row][col]
            if board[row][col] != -1:
                Qlist.append(queens[Qindex])
                break

        # Lower Left
        row = king[0]
        col = king[1]
        while row < 7 and col > 0:
            row += 1
            col -= 1
            Qindex = board[row][col]
            if board[row][col] != -1:
                Qlist.append(queens[Qindex])
                break

        # Lower Right
        row = king[0]
        col = king[1]
        while row < 7 and col < 7:
            row += 1
            col += 1
            Qindex = board[row][col]
            if board[row][col] != -1:
                Qlist.append(queens[Qindex])
                break

        return Qlist


func = Solution()
queens = [[0, 1], [1, 0], [4, 0], [0, 4], [3, 3], [2, 4]]
king = [0, 0]
# Output: [[2,2],[3,4],[4,4]]
print(func.queensAttacktheKing(queens, king))

queens = [[0, 0], [1, 1], [2, 2], [3, 4], [3, 5], [4, 4], [4, 5]]
king = [3, 3]
# Output: [[2,2],[3,4],[4,4]]
print(func.queensAttacktheKing(queens, king))

queens = [[5, 6], [7, 7], [2, 1], [0, 7], [1, 6], [5, 1], [3, 7], [0, 3], [4, 0], [1, 2], [6, 3], [5, 0], [0, 4],
          [2, 2], [1, 1], [6, 4], [5, 4], [0, 0], [2, 6], [4, 5], [5, 2], [1, 4], [7, 5], [2, 3], [0, 5], [4, 2],
          [1, 0], [2, 7], [0, 1], [4, 6], [6, 1], [0, 6], [4, 3], [1, 7]]
king = [3, 4]
# Output: [[2,3],[1,4],[1,6],[3,7],[4,3],[5,4],[4,5]]
print(func.queensAttacktheKing(queens, king))
