class Solution:
    def intToRoman(self, num: int) -> str:
        numeric = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        roman = ["M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"]
        roman_str = ""
        for i in range(0, len(numeric)):
            while num >= numeric[i]:
                num -= numeric[i]
                roman_str += roman[i]
        return roman_str

ans = Solution()
print(ans.intToRoman(96))
