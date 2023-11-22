package main

import "fmt"

// PROBLEM
// Given a string e.g "a?b?c?", replace "?" with an alphabet character a-z but the replacement cannot be the same with the
// adjacent.
//
// SAMPLE #1
// Input: a?b?c?
// Output: acbaca
//
// SAMPLE #2
// Input: k???msl?
// Output: kxyzmslp

func fillQuestionMarks(s string) string {
	arr := append([]byte{'a' - 1}, []byte(s)...)
	arr = append(arr, 'a'-1)
	for i := 1; i <= len(s); i++ {
		if arr[i] == '?' {
			for c := byte('a'); c <= 'z'; c++ {
				if arr[i-1] != c && arr[i+1] != c {
					arr[i] = c
					break
				}
			}
		}
	}
	return string(arr[1 : len(arr)-1])
}

func main() {
	fmt.Println(fillQuestionMarks("a?b?c?"))
	fmt.Println(fillQuestionMarks("k???msl?"))
}
