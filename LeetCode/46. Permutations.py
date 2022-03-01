from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        if not nums:
            return []
        used = {}
        for val in nums:
            used[val] = False

        output = []

        def dfs(remain: int, sequence: List[int]):
            if remain == 0:
                output.append(sequence)
                return

            for value in nums:
                if not used[value]:
                    used[value] = True
                    dfs(remain - 1, sequence + [value])
                    used[value] = False

        dfs(len(nums), [])
        return output


func = Solution()
print(func.permute([1, 2, 3]))
