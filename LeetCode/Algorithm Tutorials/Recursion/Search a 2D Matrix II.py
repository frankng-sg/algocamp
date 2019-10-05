class Solution:

    @staticmethod
    def binary_search(target, matrix, row=None, col=None):
        if row is None:
            low = 0
            high = len(matrix) - 1
            while low <= high:
                mid = (low + high) // 2
                if target == matrix[mid][col]:
                    return True
                if target < matrix[mid][col]:
                    high = mid - 1
                if target > matrix[mid][col]:
                    low = mid + 1
            return False
        if col is None:
            low = 0
            high = len(matrix[0]) - 1
            while low <= high:
                mid = (low + high) // 2

                if target == matrix[row][mid]:
                    return True
                if target < matrix[row][mid]:
                    high = mid - 1
                if target > matrix[row][mid]:
                    low = mid + 1
            return False

        return False

    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix:
            return False

        search_row = len(matrix) - 1
        search_col = len(matrix[0]) - 1
        while search_row >= 0 and search_col >= 0:
            if matrix[search_row][search_col] is target:
                return True
            if self.binary_search(target, matrix, row=search_row):
                return True
            if self.binary_search(target, matrix, col=search_col):
                return True
            search_row -= 1
            search_col -= 1
        return False


func = Solution()
print(func.searchMatrix(
    [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22], [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]], 20))
