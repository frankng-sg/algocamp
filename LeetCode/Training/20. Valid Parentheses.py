# 20. Valid Parentheses
# Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
#
# An input string is valid if:
#
# Open brackets must be closed by the same type of brackets.
# Open brackets must be closed in the correct order.
# Note that an empty string is also considered valid.
#
# Input: "()"
# Output: true
#
# Input: "(]"
# Output: false
#
# Input: "([)]"
# Output: false
#
# Input: "{[]}"
# Output: true

class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        mappings = {
            ')': '(',
            ']': '[',
            '}': '{'
        }
        for ch in s:
            if ch in '([{':
                stack.append(ch)
            else:
                if not stack or stack[-1] != mappings[ch]:
                    return False
                else:
                    stack.pop()

        return not stack


func = Solution()
s = ''
print(func.isValid(s))  # Output: True

s = "([)]"
print(func.isValid(s))  # Output: True

s = "{[]}"
print(func.isValid(s))  # Output: True