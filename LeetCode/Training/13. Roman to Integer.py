class Solution:
    def romanToInt(self, roman_num):
        rom2int = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
        int_num = rom2int[roman_num[0]]
        for i in range(1, len(roman_num)):
            if rom2int[roman_num[i]] <= rom2int[roman_num[i-1]]:
                int_num += rom2int[roman_num[i]]
            else:
                int_num += rom2int[roman_num[i]] - 2 * rom2int[roman_num[i-1]]
        return int_num