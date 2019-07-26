import sys, math

DEBUG = True
fact = math.factorial

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
  
n_Valid_Numbers = (K-1) * pow(K, N-1)
n_Non_Zero = 1
while (N >= n_Non_Zero + 2):
    n_Valid_Numbers -= pow(K-1, n_Non_Zero) * pow(K, N-n_Non_Zero-2)
    n_Non_Zero = n_Non_Zero + 1

        
print(n_Valid_Numbers % M)


