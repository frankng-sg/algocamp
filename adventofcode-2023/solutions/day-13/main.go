package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

type Pattern [][]byte

func readInput(filepath string) <-chan Pattern {
    ch := make(chan Pattern)
    go func() {
        defer close(ch)

        f, err := os.Open(filepath)
        if err != nil {
            fmt.Println(err)
            return
        }
        defer f.Close()

        reader := bufio.NewReader(f)
        pattern := make(Pattern, 0)
        for {
            line, err := reader.ReadString('\n')
            if err == io.EOF {
                ch <- pattern
                break
            }
            if err != nil {
                fmt.Println(err)
                return
            }
            line = strings.TrimRight(line, "\n")
            line = strings.TrimRight(line, "\r")
            if line == "" {
                ch <- pattern
                pattern = make(Pattern, 0)
            } else {
                pattern = append(pattern, []byte(line))
            }
        }
    }()
    return ch
}

func part1(filepath string, solverFn func(*Pattern) int, verbose bool) {
    fmt.Printf("\nPart 1:\n------------------------\n")
    mainBegin := time.Now()
    result := 0
    patternIndex := 0
    for pattern := range readInput(filepath) {
        patternIndex++
        subBegin := time.Now()
        cnt := solverFn(&pattern)
        result += cnt
        if verbose {
            fmt.Printf("Pattern=%d, Count=%d, Time=%s\n", patternIndex, cnt, time.Since(subBegin))
        }
    }
    fmt.Printf("\nAnswer: %d\n", result)
    if verbose {
        fmt.Printf("Total Time: %s\n", time.Since(mainBegin))
    }
}

func part2(filepath string, solverFn func(*Pattern) int, verbose bool) {
    fmt.Printf("\nPart 2:\n------------------------\n")
    mainBegin := time.Now()
    result := 0
    patternIndex := 0
    for pattern := range readInput(filepath) {
        patternIndex++
        subBegin := time.Now()
        cnt := solverFn(&pattern)
        result += cnt
        if verbose {
            fmt.Printf("Pattern=%d, Count=%d, Time=%s\n", patternIndex, cnt, time.Since(subBegin))
        }
    }
    fmt.Printf("\nAnswer: %d\n", result)
    if verbose {
        fmt.Printf("Total Time: %s\n", time.Since(mainBegin))
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ./main <input_file>")
        return
    }
    filepath := os.Args[1]
    part1(filepath, solve, true)

}
