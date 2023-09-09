package main

import "fmt"

// Time: O(n^2), Space: O(1)
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	maxlen, maxstart := 1, 0
	halflen := maxlen >> 1
	expandPalStr := func(left, right int) (strstart, strlen int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		left, right = left+1, right-1
		return left, right - left + 1
	}
	for i := 0; i < len(s)-halflen; i++ {
		var strstart, strlen int

		strstart, strlen = expandPalStr(i-1, i)
		if strlen > maxlen {
			maxlen, halflen, maxstart = strlen, strlen>>1, strstart
		}
		strstart, strlen = expandPalStr(i, i+1)
		if strlen > maxlen {
			maxlen, halflen, maxstart = strlen, strlen>>1, strstart
		}
		strstart, strlen = expandPalStr(i, i)
		if strlen > maxlen {
			maxlen, halflen, maxstart = strlen, strlen>>1, strstart
		}
	}
	return s[maxstart : maxstart+maxlen]
}

// Time: O(n^3), Space: O(1)
func longestPalindrome1(s string) string {
	if s == "" {
		return ""
	}
	isPalindrome := func(i, j int) bool {
		for ; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	maxLen, maxStart := 1, 0
	for i := 0; i <= len(s)-maxLen; i++ {
		for j := len(s) - 1; j >= i+maxLen; j-- {
			if isPalindrome(i, j) {
				maxLen = j - i + 1
				maxStart = i
			}
		}
	}
	return s[maxStart : maxStart+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))                                     // Output: "bb"
	fmt.Println(longestPalindrome("abccbaccbaabccbaabaccaccabaccccbaabaccba")) // Output: "baabccbaab"
	fmt.Println(longestPalindrome(""))                                         // Output: ""
	fmt.Println(longestPalindrome("abcba"))                                    // Output: "abcba"
	fmt.Println(longestPalindrome("abccbaabaccba"))                            // Output: "abccba"
}
