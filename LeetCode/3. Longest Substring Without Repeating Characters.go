package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	exist := make(map[byte]int)
	slen := len(s)
	max := 0
	sublen := 0
	l := 0
	for r := l; r < slen; r++ {
		if k, ok := exist[s[r]]; ok {
			for j := l; j <= k; j++ {
				delete(exist, s[j])
			}
			sublen = r - k
			l = k + 1
		} else {
			sublen++
		}
		exist[s[r]] = r
		if sublen > max {
			max = sublen
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
}
