package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count [26]int
	l := len(s)
	for i := 0; i < l; i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if count[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // Output: true
	fmt.Println(isAnagram("rat", "car"))         // Output: false
}
