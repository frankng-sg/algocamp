import sys, math

DEBUG = True
fact = math.factorial

def nCr(n, r):
  if (n < r):
    return 0
  else:    
    return fact(n) / (fact(r) * fact(n-r))

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
  
if (M >= pow(10, N)):
  valid_numbers = (K-1) * pow(K, N-1)
  if (N > 2):
    valid_numbers = valid_numbers - pow(K, N-3) * (N-2)
elif (M < pow(10, N-1)):
  valid_numbers = 0
else:
  valid_numbers = 0
    
print(valid_numbers)
