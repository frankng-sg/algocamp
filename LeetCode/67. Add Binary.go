package main

import "fmt"

func addBinary(a string, b string) string {
	result := ""
	var c, da, db byte
	for i, j := len(a)-1, len(b)-1; ; {
		if i >= 0 {
			da = a[i] - '0'
		} else {
			da = 0
		}
		if j >= 0 {
			db = b[j] - '0'
		} else {
			db = 0
		}
		sum := c + da + db
		if sum == 0 && i < 0 && j < 0 {
			break
		}
		result = string(sum&1+'0') + result
		c = sum >> 1
		i--
		j--
	}
	if result == "" {
		result = "0"
	}
	return result
}

func main() {
	fmt.Println(addBinary("100", "110010")) // Output: "0"
	fmt.Println(addBinary("0", "0"))        // Output: "0"
	fmt.Println(addBinary("0", "1"))        // Output: "1"
	fmt.Println(addBinary("1", "1"))        // Output: "10"
	fmt.Println(addBinary("11", "1"))       // Output: "100"
	fmt.Println(addBinary("1010", "1011"))  // Output: "10101"
}
