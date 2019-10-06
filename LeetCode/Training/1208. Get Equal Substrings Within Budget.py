# 1208. Get Equal Substrings Within Budget

# You are given two strings s and t of the same length. You want to change s to t.
# Changing the i-th character of s to i-th character of t costs |s[i] - t[i]| that is,
# the absolute difference between the ASCII values of the characters.
# You are also given an integer maxCost. Return the maximum length of a substring of s
# that can be changed to be the same as the corresponding substring of twith a cost less than or equal to maxCost.
# If there is no substring from s that can be changed to its corresponding substring from t, return 0.


class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:
        string_length = len(s)
        diff = [abs(ord(s[i]) - ord(t[i])) for i in range(string_length)]

        max_substring_length = 0
        spent = 0
        left = 0
        for right in range(string_length):
            spent += diff[right]
            while spent > maxCost:
                spent -= diff[left]
                left += 1
            max_substring_length = max(max_substring_length, right - left + 1)

        return max_substring_length


# func = Solution()
#
# s = 'abcd'
# t = 'bcdf'
# maxCost = 3
# print(func.equalSubstring(s, t, maxCost))  # Output: 3

# s = "abcd"
# t = "acde"
# maxCost = 0
# print(func.equalSubstring(s, t, maxCost))  # Output: 1
