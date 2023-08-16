package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	var try func(int, int, []byte)
	try = func(openP, closeP int, comb []byte) {
		if closeP > openP || openP > n || closeP > n {
			return
		}
		if openP == n && closeP == n {
			result = append(result, string(comb))
		}
		try(openP+1, closeP, append(comb, '('))
		try(openP, closeP+1, append(comb, ')'))
	}
	try(1, 0, []byte{'('})
	return result
}

func main() {
	fmt.Println(generateParenthesis(3)) // Output: ["((()))","(()())","(())()","()(())","()()()"]
}
