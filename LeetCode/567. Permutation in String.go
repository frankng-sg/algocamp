package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	target := make(map[byte]int)
	count := make(map[byte]int)
	n1 := len(s1)
	n2 := len(s2)
	if n2 < n1 {
		return false
	}
	for i := 0; i < n1; i++ {
		target[s1[i]]++
		count[s2[i]]++
	}
	diff := 0
	for i := byte('a'); i <= byte('z'); i++ {
		if target[i]-count[i] > 0 {
			diff += target[i] - count[i]
		}
	}
	l := 0
	for r := n1 - 1; r < n2; {
		if diff == 0 {
			return true
		}
		count[s2[l]]--
		if count[s2[l]] < target[s2[l]] {
			diff += 1
		}
		l++
		r++
		if r < n2 {
			if count[s2[r]] < target[s2[r]] {
				diff -= 1
			}
			count[s2[r]]++
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("axb", "abaax"))   // Output: false
	fmt.Println(checkInclusion("ab", "eidbaooo")) // Output: true
	fmt.Println(checkInclusion("axb", "aaababx")) // Output: true
}
