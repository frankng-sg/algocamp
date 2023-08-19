package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	target := [26]int{}
	count := [26]int{}
	n1 := len(s1)
	n2 := len(s2)
	if n2 < n1 {
		return false
	}
	for i := 0; i < n1; i++ {
		target[s1[i]-'a']++
		count[s2[i]-'a']++
	}
	diff := 0
	for i := 0; i < 26; i++ {
		if target[i]-count[i] > 0 {
			diff += target[i] - count[i]
		}
	}
	l := 0
	for r := n1 - 1; r < n2; {
		if diff == 0 {
			return true
		}
		count[s2[l]-'a']--
		if count[s2[l]-'a'] < target[s2[l]-'a'] {
			diff += 1
		}
		l++
		r++
		if r < n2 {
			if count[s2[r]-'a'] < target[s2[r]-'a'] {
				diff -= 1
			}
			count[s2[r]-'a']++
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("axb", "abaax"))   // Output: false
	fmt.Println(checkInclusion("ab", "eidbaooo")) // Output: true
	fmt.Println(checkInclusion("axb", "aaababx")) // Output: true
}
