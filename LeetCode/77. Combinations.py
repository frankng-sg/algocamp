# 77. Combinations
# Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
class Solution:
  def __init__(self):
    self.used = set()
    self.result = []

  def recur_search(self, n, k, depth, combination):
    if depth == k:
      self.result.append(combination)
      return
    for i in range(0, n):
      if i not in self.used:
        self.used.add(i)
        self.recur_search(i, k, depth + 1, combination + [i + 1])
        self.used.remove(i)

  def combine(self, n: int, k: int):

    self.recur_search(n, k, 0, [])
    return self.result
