package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	l := 0
	r := 0
	n := len(s)
	maxc := 0
	for ; r < n; r++ {
		count[s[r]]++
		maxc = max(maxc, count[s[r]])
		if (r-l+1)-maxc > k {
			count[s[l]]--
			l++
		}
	}
	return r - l
}

func main() {
	fmt.Println(characterReplacement("ABAB", 2))                           // Output: 4
	fmt.Println(characterReplacement("AABABBA", 1))                        // Output: 4
	fmt.Println(characterReplacement("xAABBAABA", 2))                      // Output: 6
	fmt.Println(characterReplacement("BAAAB", 2))                          // Output: 5
	fmt.Println(characterReplacement("AAAxB", 2))                          // Output: 5
	fmt.Println(characterReplacement("AAAAAAAAAabcdefghijklmnopAAAxB", 2)) // Output: 11
}
