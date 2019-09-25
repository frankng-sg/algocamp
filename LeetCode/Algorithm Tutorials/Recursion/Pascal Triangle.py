class Solution:
    def generate(self, numRows, pascal_matrix=[]):
        if numRows < len(pascal_matrix):
            return pascal_matrix

        elif numRows == 1:
            pascal_matrix = [[1]]
            return pascal_matrix
        elif numRows == 0:
            return []

        pascal_matrix = self.generate(numRows-1)

        currRow = [1]
        prevRow = pascal_matrix[numRows-2]

        for i in range(1, len(prevRow)):
            currRow.append(prevRow[i] + prevRow[i-1])
        currRow.append(1)

        pascal_matrix.append(currRow)

        return pascal_matrix
        
