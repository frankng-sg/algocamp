class Solution:
    def isPalindrome(self, s: str):
        str_len = len(s)
        for index in range(0, str_len // 2):
            if s[index] != s[str_len - 1 - index]:
                return False, index, str_len - 1 - index
        return True, 0, 0

    def validPalindrome(self, s: str) -> bool:
        palindrome, left, right = self.isPalindrome(s)
        if not palindrome:
            palindrome, unused, unused = self.isPalindrome(s[left + 1: right + 1])
            if not palindrome:
                palindrome, unused, unused = self.isPalindrome(s[left: right])
                return palindrome
        return True


answer = Solution()

print(answer.validPalindrome("abc"))
print(answer.validPalindrome("d____"))
print(answer.validPalindrome("_d___"))
print(answer.validPalindrome("__d__"))
print(answer.validPalindrome("___d_"))
print(answer.validPalindrome("____d"))

