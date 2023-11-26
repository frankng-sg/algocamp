package main

import (
    "fmt"
)

func reverseWords(s string) string {
    var out, word string
    s = s + " "
    for _, c := range s {
        if c != ' ' {
            word += string(c)
        } else if word != "" {
            out = word + " " + out
            word = ""
        }
    }
    return out[:len(out)-1]
}

func main() {
    fmt.Printf("'%s'\n", reverseWords("abc 123 xyz"))
}
