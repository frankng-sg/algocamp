package main

import "fmt"

func minWindow(s string, t string) string {
	target := make(map[byte]int)
	count := make(map[byte]int)
	for i := range t {
		target[t[i]] += 1
	}
	flag := len(t)
	slen := len(s)
	minLen := len(s) + len(t)
	minStr := ""
	l, r := 0, 0
	for l <= r {
		for ; flag > 0 && r < slen; r++ {
			if count[s[r]] < target[s[r]] {
				flag--
			}
			count[s[r]]++
		}
		if flag > 0 {
			return minStr
		}
		if r-l < minLen {
			minLen = r - l
			minStr = s[l:r]
		}
		count[s[l]]--
		if count[s[l]] < target[s[l]] {
			flag++
		}
		l++
	}
	return minStr
}

func main() {
	fmt.Printf("\"%s\"\n", minWindow("ADOBECODEBANC", "ABC"))          // Output: "BANC"
	fmt.Printf("\"%s\"\n", minWindow("ab", "a"))                       // Output: "a"
	fmt.Printf("\"%s\"\n", minWindow("a", "a"))                        // Output: "a"
	fmt.Printf("\"%s\"\n", minWindow("aDOBECODEBANC", "ABC"))          // Output: "BANC"
	fmt.Printf("\"%s\"\n", minWindow("a", "aa"))                       // Output: ""
	fmt.Printf("\"%s\"\n", minWindow("xxxxAxAxBxBxCxCAABB", "AABBCC")) // Output: "CxCAABB"
}
