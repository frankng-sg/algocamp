# 1218. Longest Arithmetic Subsequence of Given Difference
# Given an integer array arr and an integer difference, return the length of the longest subsequence in arr which is
# an arithmetic sequence such that the difference between adjacent elements in the subsequence equals difference.
#
# Example 1:
# Input: arr = [1,2,3,4], difference = 1
# Output: 4
# Explanation: The longest arithmetic subsequence is [1,2,3,4].

# Example 2:
# Input: arr = [1,3,5,7], difference = 1
# Output: 1
# Explanation: The longest arithmetic subsequence is any single element.

# Example 3:
# Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
# Output: 4
# Explanation: The longest arithmetic subsequence is [7,5,3,1].

class Solution:
    def longestSubsequence(self, arr, difference):
        lookup = {}
        longest = 1
        for i in range(0, len(arr)):
            if arr[i] - difference not in lookup:
                lookup[arr[i]] = 1
            else:
                lookup[arr[i]] = lookup[arr[i] - difference] + 1
                longest = max(lookup[arr[i]], longest)
        return longest


func = Solution()
arr = [1, 5, 7, 8, 5, 3, 4, 2, 1]
difference = -2
print(func.longestSubsequence(arr, difference)) # Output: 4
