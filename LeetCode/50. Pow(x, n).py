class Solution:
    def myPow(self, x, n):

        def power_of(num,power, value = 1):
            if power == 0:
              return value

            if power == 1:
                return value * num

            elif power % 2 != 0:
                return power_of(num*num, (power - 1) // 2, value * num)

            else:
                return power_of(num*num, power // 2, value)

        if (x == 0):
            return 0

        if (n < 0):
            return 1.0 / power_of(x, -n)

        else:
            return power_of(x, n)
