package main

import (
	"fmt"
	"sync"
)

func groupAnagramsWithGoroutines(strs []string) [][]string {
	m := make(map[[26]int][]string)
	var mLock sync.Mutex
	var wg sync.WaitGroup
	for _, s := range strs {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			var count [26]int
			for _, ch := range str {
				count[ch-'a']++
			}
			mLock.Lock()
			m[count] = append(m[count], str)
			mLock.Unlock()
		}(s)
	}
	wg.Wait()
	result := make([][]string, len(m))
	i := 0
	for _, r := range m {
		result[i] = r
		i++
	}
	return result
}

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
