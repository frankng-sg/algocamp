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

func groupAnagrams(strs []string) [][]string {
	used := make(map[int]bool)
	l := len(strs)
	var result [][]string
	for i := 0; i < l; i++ {
		if _, ok := used[i]; !ok {
			used[i] = true
			group := []string{strs[i]}
			for j := i + 1; j < l; j++ {
				if _, ok := used[j]; !ok && isAnagram(strs[i], strs[j]) {
					used[j] = true
					group = append(group, strs[j])
				}
			}
			result = append(result, group)
		}
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams([]string{""}))                                       // Output: [[""]]
	fmt.Println(groupAnagrams([]string{"a"}))                                      // Output: [["a"]]
}
