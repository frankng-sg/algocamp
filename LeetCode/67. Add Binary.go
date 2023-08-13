package main

import "fmt"

func higherInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	lm := higherInt(la, lb)
	result := make([]byte, lm+1)
	var c byte
	for i, j, k := la-1, lb-1, lm; k >= 0; k-- {
		sum := c
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		result[k] = sum&1 + '0'
		c = sum >> 1
	}
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}

func main() {
	fmt.Println(addBinary("100", "110010")) // Output: "110110"
	fmt.Println(addBinary("0", "0"))        // Output: "0"
	fmt.Println(addBinary("0", "1"))        // Output: "1"
	fmt.Println(addBinary("1", "1"))        // Output: "10"
	fmt.Println(addBinary("11", "1"))       // Output: "100"
	fmt.Println(addBinary("1010", "1011"))  // Output: "10101"
}
