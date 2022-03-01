from typing import List


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if digits == "":
            return []

        keypads = [None, None, 'abc', 'def', 'ghi', 'jkl', 'mno', 'pqrs', 'tuv', 'wxyz']
        output = []
        size = len(digits)

        def dfs(loc, sequence):
            if loc == size:
                output.append(sequence)
                return

            num = int(digits[loc])
            for letter in keypads[num]:
                dfs(loc + 1, sequence + letter)

        dfs(0, "")

        return output


func = Solution()
print(func.letterCombinations("23"))
