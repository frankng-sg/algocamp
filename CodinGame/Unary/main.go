package main

import (
    "fmt"
    "strings"
)
import "os"
import "bufio"

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    scanner.Scan()
    MESSAGE := scanner.Text()

    prev := 0
    count := 0
    var out strings.Builder

    write := func() {
        if count == 0 {
            return
        }
        if prev == 0 {
            out.WriteString(" 00 " + strings.Repeat("0", count))
        } else {
            out.WriteString(" 0 " + strings.Repeat("0", count))
        }
    }
    for _, c := range MESSAGE {
        for i := 6; i >= 0; i-- {
            b := int((c >> i) & 1)
            if b != prev {
                write()
                count = 1
                prev = b
            } else {
                count++
            }
        }
    }
    write()
    fmt.Println(out.String()[1:])
}
