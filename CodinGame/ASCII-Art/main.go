package main

import (
	"fmt"
)
import "os"
import "bufio"

type AsciiArt []string

var text [27]AsciiArt

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	for i := 0; i < 27; i++ {
		text[i] = make(AsciiArt, H)
	}

	scanner.Scan()
	T := scanner.Text()

	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		for j := 0; j < 27; j++ {
			text[j][i] = ROW[j*L : (j+1)*L]
		}
	}

	for i := 0; i < H; i++ {
		for _, c := range T {
			var j int32
			if c >= 'a' && c <= 'z' {
				j = c - 'a'
			} else if c >= 'A' && c <= 'Z' {
				j = c - 'A'
			} else {
				j = 26
			}
			fmt.Print(text[j][i])
		}
		fmt.Println()
	}
}
