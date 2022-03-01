class Solution:
    def power_value(self, num):
        power = 0
        while num != 1:
            if num % 2 == 0:
                num = num // 2
            else:
                num = 3 * num + 1
            power += 1
        return power

    def getKth(self, lo: int, hi: int, k: int) -> int:
        numbers = []
        for i in range(lo, hi + 1):
            numbers.append([i, self.power_value(i)])

        numbers.sort(key=lambda x: x[1])
        return numbers[k - 1][0]


ans = Solution()
print(ans.getKth(12, 15, 2))  # 13
print(ans.getKth(1, 1, 1))  # 1
print(ans.getKth(7, 11, 4))  # 7
print(ans.getKth(10, 20, 5))  # 13
print(ans.getKth(1, 1000, 777))  # 570
