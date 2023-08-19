package main

import "fmt"

func reverseBits(num uint32) uint32 {
	rev := uint32(0)
	for i := 0; i < 32; i++ {
		rev = (rev << 1) + (num & 1)
		num >>= 1
	}
	return rev
}

func main() {
	fmt.Println(reverseBits(0))       // Output: 0
	fmt.Println(reverseBits(1 << 31)) // Output: 1
}
