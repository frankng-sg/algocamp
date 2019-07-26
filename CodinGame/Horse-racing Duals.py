import sys
import math

# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = int(input())
strength = []
for i in range(n):
    strength.append(int(input()))
    
strength.sort()

min_diff = strength[1] - strength[0]
for i in range(2, n):
    if (strength[i] - strength[i-1] < min_diff):
        min_diff = strength[i] - strength[i-1]
    
print(min_diff)