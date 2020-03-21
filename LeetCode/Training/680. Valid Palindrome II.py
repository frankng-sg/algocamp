class Solution:
    def validPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1
        while left < right:
            if s[left] == s[right]:
                left, right = left + 1, right - 1
            else:
                case1 = s[left: right]
                case2 = s[left + 1: right + 1]
                return case1 == case1[::-1] or case2 == case2[::-1]
        return True


answer = Solution()

print(answer.validPalindrome("abc"))
print(answer.validPalindrome("d____"))
print(answer.validPalindrome("_d___"))
print(answer.validPalindrome("__d__"))
print(answer.validPalindrome("___d_"))
print(answer.validPalindrome("____d"))

