package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

const (
    MAX_SIZE     = 1000
    SCALE_FACTOR = 1000000
)

func readInput(filepath string) Universe {
    f, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    u := make(Universe, 0, MAX_SIZE)
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        line = strings.TrimRight(line, "\n")
        line = strings.TrimRight(line, "\r")
        if len(line) == 0 {
            break
        }
        row := make([]Space, len(line))
        for i := len(line) - 1; i >= 0; i-- {
            row[i] = ToSpace(line[i])
        }
        u = append(u, row)
    }
    return u
}

func solvePart1(u Universe) {
    u.Expand()
    nodes := u.Find(Galaxy)
    sumDist := 0
    for i := range nodes {
        for j := i + 1; j < len(nodes); j++ {
            sumDist += Distance(nodes[i], nodes[j])
        }
    }
    fmt.Printf("\nPart 1: %d\n", sumDist)
}

func solvePart2(u Universe) {
    nodes := u.Find(Galaxy)
    emptyCols := u.FindEmptyCols()
    emptyRows := u.FindEmptyRows()
    sumDist := 0
    for i := range nodes {
        for j := i + 1; j < len(nodes); j++ {
            nRows := emptyRows.Count(nodes[i].Row, nodes[j].Row)
            if nRows > 0 {
                nRows = nRows * (SCALE_FACTOR - 1)
            }
            nCols := emptyCols.Count(nodes[i].Col, nodes[j].Col)
            if nCols > 0 {
                nCols = nCols * (SCALE_FACTOR - 1)
            }
            sumDist += Distance(nodes[i], nodes[j]) + nRows + nCols

        }
    }
    fmt.Printf("\nPart 2: %d\n", sumDist)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <input_file>")
        return
    }
    filepath := os.Args[1]
    u := readInput(filepath)

    solvePart2(u)

    solvePart1(u)
}
