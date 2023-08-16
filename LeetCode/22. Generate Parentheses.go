package main

import "fmt"

func generateParenthesis(n int) []string {
	openP := 0
	closeP := 0
	lim := n << 1
	comb := make([]byte, lim)
	result := []string{}
	var try func(int)
	try = func(k int) {
		if k == lim {
			result = append(result, string(comb))
		}
		if openP < n {
			comb[k] = '('
			openP++
			try(k + 1)
			openP--
		}
		if closeP < openP && closeP < n {
			comb[k] = ')'
			closeP++
			try(k + 1)
			closeP--
		}
	}
	try(0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3)) // Output: ["((()))","(()())","(())()","()(())","()()()"]
}
