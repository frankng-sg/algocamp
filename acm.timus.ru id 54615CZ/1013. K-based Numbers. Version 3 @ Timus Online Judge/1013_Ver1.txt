import sys, math

DEBUG = True
DEBUG_MSG = "-----DEBUG MESSAGE-----\n"
fact = math.factorial

def valid_Kbased_Ndigits(K, N):
    if (N <= 0):
        return 0
    n_Valid_Numbers = (K-1) * pow(K, N-1)
    n_Non_Zero = 1
    while (N >= n_Non_Zero + 2):
        n_Valid_Numbers -= (K-1) * pow(K, N-n_Non_Zero-2)
        n_Non_Zero = n_Non_Zero + 1
    return n_Valid_Numbers

if DEBUG:
  try:
    fp = open('1013.txt', 'r')
    N = fp.readline()
    K = fp.readline()
    M = fp.readline()
  finally:
    fp.close()
else:
  N = sys.stdin.readline()
  K = sys.stdin.readline()
  M = sys.stdin.readline()

N = int(N)
K = int(K)
M = int(M)
  
n_Valid_Numbers = 0

if (M <= pow(K, N-1)):
    n_Valid_Numbers = 0
    DEBUG_MSG += "M <= K^(N-1)\n"
elif (M >= pow(K, N)):
    n_Valid_Numbers = valid_Kbased_Ndigits(K, N)
    DEBUG_MSG += "M >= K^N\n"
else:
    DEBUG_MSG += "K^(N-1) < M < K^N\n"
    n_Valid_Numbers = 0
        
print(n_Valid_Numbers)

if DEBUG:
    print(DEBUG_MSG)

