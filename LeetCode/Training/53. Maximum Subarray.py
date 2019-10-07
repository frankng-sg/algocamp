# 53. Maximum Subarray
# Given an integer array nums, find the contiguous subarray (containing at least one number) which
# has the largest sum and return its sum.
#
# Input: [-2,1,-3,4,-1,2,1,-5,4],
# Output: 6
# Explanation: [4,-1,2,1] has the largest sum = 6.
class Solution:
    def find_max_cross_sum(self, nums, low, mid, high):
        max_sum = sum = nums[mid]
        for i in range(mid + 1, high + 1):
            sum += nums[i]
            if max_sum < sum:
                max_sum = sum

        sum = max_sum
        for i in range(mid - 1, -1, -1):
            sum += nums[i]
            if max_sum < sum:
                max_sum = sum

        return max_sum

    def recur_search(self, nums, low, high):
        if low == high:
            return nums[low]

        mid = (low + high) // 2
        left_sum = self.recur_search(nums, low, mid)
        right_sum = self.recur_search(nums, mid + 1, high)
        cross_sum = self.find_max_cross_sum(nums, low, mid, high)
        return max(left_sum, right_sum, cross_sum)

    def maxSubArray(self, nums):
        return self.recur_search(nums, 0, len(nums) - 1)


# func = Solution()
#
# nums = [0, 5, 0, 0, 0, 1, 2, 3, 4]
# print(func.maxSubArray(nums))  # Output: 15

# nums = [-2, -7, -3, -4, -10, -2, -1, -5, -4]
# print(func.maxSubArray(nums))  # Output: -1

# nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
# print(func.maxSubArray(nums))  # Output: 6

# nums = [2,0,3,-2]
# print(func.maxSubArray(nums))  # Output: 5