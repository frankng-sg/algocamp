def longestPalindrome(s: str) -> str:
    palindrome = {}

    def isPalindrome(start_pos, end_pos):
        if start_pos >= end_pos:
            return True

        if (start_pos, end_pos) in palindrome:
            return True
    
        if isPalindrome(start_pos + 1, end_pos - 1) and s[start_pos] == s[end_pos]:
            palindrome[start_pos, end_pos] = True
            return True

        return False

    for i in range(len(s)):
        palindrome[(i, i)] = True

    max_len = 1
    max_start = 0
    max_end = 0

    for end_pos in range(len(s) - 1, 0, -1):        
        for start_pos in range(0, end_pos - max_len + 1):
            if isPalindrome(start_pos, end_pos):
                max_len = end_pos - start_pos + 1
                max_start = start_pos
                max_end = end_pos
                print(start_pos, end_pos, max_len)
                break

    
    return s[max_start:max_end + 1]


#print(longestPalindrome('aa'))
#print(longestPalindrome('ababababababababa'))
#print(longestPalindrome('bananas'))
print(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))