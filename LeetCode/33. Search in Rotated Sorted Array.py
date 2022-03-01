class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """

        def dfs(left, right):
            if left > right:
                return -1

            mid = (left + right) // 2
            if nums[mid] == target:
                return mid

            if nums[mid] < target:
                if nums[left] > nums[mid]:
                    location = dfs(left, mid - 1)
                    if location != -1:
                        return location
                return dfs(mid + 1, right)
            else:
                if nums[mid] > nums[right]:
                    location = dfs(mid + 1, right)
                    if location != -1:
                        return location
                return dfs(left, mid - 1)

        return dfs(0, len(nums) - 1)


func = Solution()
nums = [4, 5, 6, 7, 0, 1, 2]
target = 0
print(func.search(nums, target))

nums = [4, 5, 6, 7, 0, 1, 2]
target = 3
print(func.search(nums, target))