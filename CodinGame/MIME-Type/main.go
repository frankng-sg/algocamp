package main

import (
    "fmt"
    "strings"
)
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 1000000), 1000000)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &N)

    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(), &Q)

    m := make(map[string]string)

    for i := 0; i < N; i++ {
        // EXT: file extension
        // MT: MIME type.
        var EXT, MT string
        scanner.Scan()
        fmt.Sscan(scanner.Text(), &EXT, &MT)
        m[strings.ToUpper(EXT)] = MT
    }
    for i := 0; i < Q; i++ {
        scanner.Scan()
        FNAME := scanner.Text()
        tmp := strings.Split(FNAME, ".")
        if len(tmp) <= 1 {
            fmt.Println("UNKNOWN")
            continue
        }
        ext := strings.ToUpper(tmp[len(tmp)-1])
        if v, ok := m[ext]; !ok {
            fmt.Println("UNKNOWN")
        } else {
            fmt.Println(v)
        }
    }
}
