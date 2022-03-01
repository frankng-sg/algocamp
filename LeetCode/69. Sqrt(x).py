class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 1:
            return 1

        left, right = 0, x // 2
        while left <= right:
            mid = (left + right) // 2
            if mid * mid == x:
                return mid
            elif mid * mid < x:
                left = mid + 1
            else:
                right = mid - 1

        return right

func = Solution()
print(func.mySqrt(0))
print(func.mySqrt(1))
print(func.mySqrt(2))
print(func.mySqrt(3))
print(func.mySqrt(8))
print(func.mySqrt(15))
print(func.mySqrt(16))