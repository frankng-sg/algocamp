package main

import "fmt"

// Time: O(n^3), Space: O(1)
func longestPalindrome(s string) string {
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
