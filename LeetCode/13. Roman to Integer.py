class Solution:
    def romanToInt(self, roman_num):
        roman2int = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
        num_len = len(roman_num)
        result = roman2int[roman_num[num_len - 1]]
        for i in range(num_len - 1, 0, -1):
            if roman2int[roman_num[i]] > roman2int[roman_num[i - 1]]:
                result -= roman2int[roman_num[i - 1]]
            else:
                result += roman2int[roman_num[i - 1]]

        return result

ans = Solution()

print(ans.romanToInt("III"))  # 3
print(ans.romanToInt("XCVI")) # 96
print(ans.romanToInt("IV"))  # 4