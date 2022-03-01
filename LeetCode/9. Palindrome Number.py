class Solution:
    def isPalindrome(self, x):
        if x < 0:
            return False
        
        if x < 10:
            return True
        
        if x % 10 == 0:
            return False
                    
        x_reversed = 0
        while True:
            x_reversed = int(x_reversed * 10 + (x % 10))
            x = int(x / 10)
                      
            if x <= x_reversed:
                break
                
        if x_reversed < 10:
            return x == x_reversed
        else:
            return x == x_reversed or x == int(x_reversed / 10)