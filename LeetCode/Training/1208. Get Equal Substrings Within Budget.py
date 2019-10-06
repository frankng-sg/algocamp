# 1208. Get Equal Substrings Within Budget

# You are given two strings s and t of the same length. You want to change s to t.
# Changing the i-th character of s to i-th character of t costs |s[i] - t[i]| that is,
# the absolute difference between the ASCII values of the characters.
# You are also given an integer maxCost. Return the maximum length of a substring of s
# that can be changed to be the same as the corresponding substring of twith a cost less than or equal to maxCost.
# If there is no substring from s that can be changed to its corresponding substring from t, return 0.


class Solution:
    def equalSubstring(self, s: str, t: str, maxCost: int) -> int:
        i = 0
        j = 0
        spent = 0
        string_size = len(s)
        max_substring_size = 0
        while i < string_size and j < string_size:
            while j < string_size:
                cost = abs(ord(s[j]) - ord(t[j]))
                if spent + cost <= maxCost:
                    spent += cost
                    j += 1
                else:
                    break

            if max_substring_size < j - i:
                max_substring_size = j - i

            spent -= abs(ord(s[i]) - ord(t[i]))
            i += 1

        return max_substring_size


func = Solution()

# s = 'abcd'
# t = 'bcdf'
# maxCost = 3
# print(func.equalSubstring(s, t, maxCost))  # Output: 3

# s = "abcd"
# t = "acde"
# maxCost = 0
# print(func.equalSubstring(s, t, maxCost))  # Output: 1
