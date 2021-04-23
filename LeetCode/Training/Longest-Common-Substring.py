# Python3 implementation of Finding
# Length of Longest Common Substring
 
# Returns length of longest common
# substring of X[0..m-1] and Y[0..n-1]
 
 
def LCSubStr(X, Y):
 
    # Create a table to store lengths of
    # longest common suffixes of substrings.
    # Note that lcs[i][j] contains the
    # length of longest common suffix of
    # X[0...i-1] and Y[0...j-1]. The first
    # row and first column entries have no
    # logical meaning, they are used only
    # for simplicity of the program.
    m = len(X)
    n = len(Y)


    # lcs is the table with zero
    # value initially in each cell
    lcs = [[0 for k in range(n+1)] for l in range(m+1)]
    prev = [[-1 for k in range(n+1)] for l in range(m+1)]
 
    # To store the length of
    # longest common substring
    max_len = 0
    max_start = 0
    max_end = 0
 
    # Following steps to build
    # lcs[m+1][n+1] in bottom up fashion
    for i in range(m + 1):
        for j in range(n + 1):
            if (i == 0 or j == 0):
                continue

            if (X[i-1] == Y[j-1]):
                if prev[i-1][j-1] == -1:
                    prev[i-1][j-1] = i-1
                prev[i][j] = prev[i-1][j-1]
                lcs[i][j] = lcs[i-1][j-1] + 1
                if max_len < lcs[i][j]:
                    max_len = lcs[i][j]
                    max_start = prev[i-1][j-1]
                    max_end = i-1
    if max_len == 0:
        return ''
    else:
        return X[max_start:max_end+1]
 
 

X = 'OldSite:GeeksforGeeks.org'
Y = 'NewSite:GeeksQuiz.com'


#X = 'abc'
#Y = 'xbxxxab'

print(LCSubStr(X, Y))