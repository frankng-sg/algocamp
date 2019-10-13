class Solution:
    def balancedStringSplit(self, s: str) -> int:
        count = 0
        balance = 0
        lookup = {'L':-1, 'R':1}
        for ch in s:
            balance += lookup[ch]
            if balance == 0:
                count += 1
        return count