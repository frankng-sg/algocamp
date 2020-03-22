INT32_UPPER_LIMIT = 2 ** 31 - 1
INT32_LOWER_LIMIT = -2 ** 31

class Solution:
    def reverse(self, x):
        # Obtain the sign of x
        negative = True if x < 0 else False
        x = -x if x < 0 else x
        rev_num = 0
        while x > 0:
            rev_num = rev_num * 10 + x % 10
            x //= 10

        if rev_num > INT32_UPPER_LIMIT:
            return 0

        return -rev_num if negative else rev_num

ans = Solution()
print(ans.reverse(103))
print(ans.reverse(-103))