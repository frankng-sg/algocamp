package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, ch := range s {
			count[ch-'a']++
		}
		m[count] = append(m[count], s)
	}
	result := make([][]string, len(m))
	i := 0
	for _, r := range m {
		result[i] = r
		i++
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})) // Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	fmt.Println(groupAnagrams([]string{""}))                                       // Output: [[""]]
	fmt.Println(groupAnagrams([]string{"a"}))                                      // Output: [["a"]]
}
