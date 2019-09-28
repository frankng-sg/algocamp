class Solution:
    def myPow(self, x, n):

        if x == 0:
            return 0

        def recur_pow(x, n):
            if n == 0:
                return 1

            if (n % 2 == 0):
                result = recur_pow(x, n / 2)
                return result * result
            else:
                result = recur_pow(x, (n - 1) / 2)
                return result * result * x

        if (n < 0):
            return 1.0 / recur_pow(x, -n)
        else:
            return recur_pow(x, n)
