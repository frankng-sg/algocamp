class Solution:
    def climbStairs(self, n):
        cache = {}

        def howManyWays(n):
            if n < 3:
                return n

            if n in cache:
                return cache[n]

            result = howManyWays(n - 1) + howManyWays(n - 2)

            cache[n] = result

            return result

        return howManyWays(n)
