class Solution:
    @staticmethod
    def recur_search(matrix, target, search_row, search_col):
        if search_col < 0 or search_row >= len(matrix):
            return False

        if matrix[search_row][search_col] == target:
            return True
        if matrix[search_row][search_col] > target:
            return Solution.recur_search(matrix, target, search_row, search_col - 1)
        return Solution.recur_search(matrix, target, search_row + 1, search_col)

    def searchMatrix(self, matrix, target):
        if not matrix:
            return False

        return Solution.recur_search(matrix, target, 0, len(matrix[0]) - 1)


func = Solution()

# Output: False
matrix = [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22], [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]]
target = 20
print(func.searchMatrix(matrix, target))

# Output: True
matrix = [[48, 65, 70, 113, 133, 163, 170, 216, 298, 389], [89, 169, 215, 222, 250, 348, 379, 426, 469, 554],
          [178, 202, 253, 294, 367, 392, 428, 434, 499, 651], [257, 276, 284, 332, 380, 470, 516, 561, 657, 698],
          [275, 331, 391, 432, 500, 595, 602, 673, 758, 783], [357, 365, 412, 450, 556, 642, 690, 752, 801, 887],
          [359, 451, 534, 609, 654, 662, 693, 766, 803, 964], [390, 484, 614, 669, 684, 711, 767, 804, 857, 1055],
          [400, 515, 683, 732, 812, 834, 880, 930, 1012, 1130], [480, 538, 695, 751, 864, 939, 966, 1027, 1089, 1224]]
target = 642
print(func.searchMatrix(matrix, target))
