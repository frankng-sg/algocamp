class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        OVERFLOW = 2**31 - 1
        # special cases
        if divisor == 0:
            return OVERFLOW
        if dividend == 0:
            return 0
        if divisor == 1:
            return dividend

        # edge case
        if divisor == -1 and dividend == -2**31:
            return OVERFLOW

        negative_result = (divisor < 0) ^ (dividend < 0)
        dividend, divisor = abs(dividend), abs(divisor)

        sum = 0
        quotient = 0

        for power in range(31, -1, -1):
            shifted = divisor << power
            # check if left-shift operation does not result in OVERFLOW
            if shifted >> power == divisor:
                # left-shift operation is valid (does not result in loss of information)
                if sum + shifted <= dividend:
                    sum += shifted
                    quotient |= 1 << power

        if negative_result:
            return -quotient
        return quotient
