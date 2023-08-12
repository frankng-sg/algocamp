package main

import (
	"fmt"
	"strings"
)

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	result := true
	for {
		for l < r && !isAlphanumeric(s[l]) {
			l++
		}
		for l < r && !isAlphanumeric(s[r]) {
			r--
		}
		if l >= r {
			break
		}
		if s[l] != s[r] {
			result = false
			break
		}
		l++
		r--
	}
	return result
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Output: true
	fmt.Println(isPalindrome("race a car"))                     // Output: false
	fmt.Println(isPalindrome("   "))                            // Output: true
}
