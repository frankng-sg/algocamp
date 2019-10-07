# 53. Maximum Subarray
# Given an integer array nums, find the contiguous subarray (containing at least one number) which
# has the largest sum and return its sum.
#
# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.
class Solution:
    def maxSubArray(self, nums):
        gain = 0
        loss = 0
        max_gain = nums[0]
        for i in range(0, len(nums)):
            if nums[i] < 0:
                loss += nums[i]
                max_gain = max(max_gain, nums[i])
            else:
                if gain + loss >= 0:
                    gain = nums[i] + loss + gain
                    loss = 0
                else:
                    gain = nums[i]
                    loss = 0
                max_gain = max(max_gain, gain)
        return max_gain

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
