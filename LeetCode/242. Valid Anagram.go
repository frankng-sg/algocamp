package main

import "fmt"

func isAnagram(s string, t string) bool {
	count := make(map[byte]int)
	slen, tlen := len(s), len(t)
	for i := 0; i < slen; i++ {
		if _, ok := count[s[i]]; ok {
			count[s[i]]++
		} else {
			count[s[i]] = 1
		}
	}
	for i := 0; i < tlen; i++ {
		if _, ok := count[t[i]]; !ok {
			return false
		}
		count[t[i]]--
		if count[t[i]] == 0 {
			delete(count, t[i])
		}
	}
	return len(count) == 0
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
