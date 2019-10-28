class Solution(object):
    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 0 or nums is None:
            return None

        if len(nums) == 1:
            return nums[0]

        left, right = 0, len(nums) - 1
        while left + 1 < right:
            mid = (left + right) // 2
            if nums[left] > nums[mid]:
                right = mid
            else:
                left = mid

        return min(nums[0], nums[left], nums[right])

func = Solution()

# Input: [3,4,5,6,7,0]
# Output: 0

nums = [3, 4, 5, 6, 7, 0]
print(func.findMin(nums))

# Input: [0,1,2,3,5,6,7]
# Output: 0

nums = [0, 1, 2, 3, 5, 6, 7]
print(func.findMin(nums))

# Input: [2, 1]
# Output: 1

nums = [2, 1]
print(func.findMin(nums))

# Input: [3,4,5,1,2]
# Output: 1

nums = [3, 4, 5, 1, 2]
print(func.findMin(nums))

# Input: [4,5,6,7,0,1,2]
# Output: 0

nums = [4, 5, 6, 7, 0, 1, 2]
print(func.findMin(nums))

# Input: [7, 0, 1, 2, 3, 4, 5, 6]
# Output: 0

nums = [7, 0, 1, 2, 3, 4, 5, 6]
print(func.findMin(nums))
