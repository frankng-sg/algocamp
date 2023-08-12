package main

import (
	"fmt"
	"strings"
)

func isAlphanumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var ns []rune
	for _, v := range s {
		if isAlphanumeric(v) {
			ns = append(ns, v)
		}
	}
	for i, j := 0, len(ns)-1; i < j; {
		if ns[i] != ns[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // Output: true
	fmt.Println(isPalindrome("race a car"))                     // Output: false
	fmt.Println(isPalindrome("   "))                            // Output: true
}
