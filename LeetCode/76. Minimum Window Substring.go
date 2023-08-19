package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	target := [84]int{}
	count := [84]int{}
	for i := range t {
		target[t[i]-'A'] += 1
	}
	flag := len(t)
	slen := len(s)
	minLen := math.MaxInt
	minStr := ""
	l, r := 0, 0
	for l <= r {
		for ; flag > 0 && r < slen; r++ {
			if count[s[r]-'A'] < target[s[r]-'A'] {
				flag--
			}
			count[s[r]-'A']++
		}
		if flag > 0 {
			return minStr
		}
		if r-l < minLen {
			minLen = r - l
			minStr = s[l:r]
		}
		count[s[l]-'A']--
		if count[s[l]-'A'] < target[s[l]-'A'] {
			flag++
		}
		l++
	}
	return minStr
}

func main() {
	fmt.Printf("\"%s\"\n", minWindow("ADOBECODEBANC", "ABC"))          // Output: "BANC"
	fmt.Printf("\"%s\"\n", minWindow("AB", "A"))                       // Output: "A"
	fmt.Printf("\"%s\"\n", minWindow("A", "A"))                        // Output: "A"
	fmt.Printf("\"%s\"\n", minWindow("ADOBECODEBANC", "ABC"))          // Output: "BANC"
	fmt.Printf("\"%s\"\n", minWindow("A", "AA"))                       // Output: ""
	fmt.Printf("\"%s\"\n", minWindow("zXXXAXAXBXBXCXCAABB", "AABBCC")) // Output: "CXCAABB"
}
