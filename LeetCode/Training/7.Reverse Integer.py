INT32_UPPER_LIMIT = 2147483647
INT32_LOWER_LIMIT = -2147483648

class Solution:
    def reverse(self, x):
        # Obtain the sign of x
        if x >= 0:
            sign = ''
        else:
            sign = '-'

        # First we convert absolute value of x to string then reverse it.
        str_x = str(abs(x))
        str_x = str_x[::-1]
        # Second we add 'sign' in front of reversed string
        str_x = sign + str_x
        # Third we check if the new reverse integer is within range
        reverse_x = int(str_x)
        if INT32_LOWER_LIMIT <= reverse_x <= INT32_UPPER_LIMIT:
            return reverse_x
        else:
            return 0
