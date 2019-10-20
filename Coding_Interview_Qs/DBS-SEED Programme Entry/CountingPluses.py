"""
Count number of pluses in a matrix

Example 1

Input
3
010
111
010

Output
1

Example 2

Input
5
00100
00100
11111
00100
00100

Output
2

Example 3

Input
8
00010000
00010000
00010000
11111111
00010000
00010010
00010111
00010010

Output
3

"""

def isPlus(matrix, i, j, size):
    if i + size // 2 >= N or j + size // 2 >= N:
        return False

    for k in range(i - (size // 2), i + (size // 2)):
        if matrix[k][j] == "0":
            return False

    for k in range(j - (size // 2), j + (size // 2)):
        if matrix[i][k] == "0":
            return False

    for row in range(i - (size // 2), i):
        for col in range(j - (size // 2), j):
            if matrix[row][col] == "1":
                return False

    for row in range(i + (size // 2), i, -1):
        for col in range(j + (size // 2), j, -1):
            if matrix[row][col] == "1":
                return False

    for row in range(i - (size // 2), i):
        for col in range(j + (size // 2), j, -1):
            if matrix[row][col] == "1":
                return False

    for row in range(i + (size // 2), i, -1):
        for col in range(j - (size // 2), j):
            if matrix[row][col] == "1":
                return False

    return True


def main():
    N = int(input())
    matrix = []
    for i in range(N):
        matrix.append(input())
    output = 0
    for i in range(1, N - 1):
        for j in range(1, N - 1):
            if matrix[i][j] == "1" \
                    and matrix[i][j - 1] == "1" \
                    and matrix[i][j + 1] == "1" \
                    and matrix[i - 1][j] == "1" \
                    and matrix[i + 1][j] == "1":
                for size in range(3, N + 1, 2):
                    if isPlus(matrix, i, j, size):
                        output += 1
                    else:
                        break
    print(output)


if __name__ == '__main__':
    main()


