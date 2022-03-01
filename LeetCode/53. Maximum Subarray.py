# 53. Maximum Subarray
# Given an integer array nums, find the contiguous subarray (containing at least one number) which
# has the largest sum and return its sum.
#
# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.


class Solution:
    def maxSubArray(self, nums):
        current_sum = 0
        max_sum = nums[0]
        for num in nums:
            current_sum = max(num, current_sum + num)
            if max_sum < current_sum:
                max_sum = current_sum

        return max_sum

# func = Solution()
#
# nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
# print(func.maxSubArray(nums))  # Output: 6
#
# nums = [0, 5, 0, 0, 0, 1, 2, 3, 4]
# print(func.maxSubArray(nums))  # Output: 15
# nums = [-2, -7, -3, -4, -10, -2, -1, -5, -4]
# print(func.maxSubArray(nums))  # Output: -1
# nums = [2,0,3,-2]
# print(func.maxSubArray(nums))  # Output: 5
