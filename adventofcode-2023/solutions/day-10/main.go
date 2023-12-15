package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "os"
)

const (
    MAX_SIZE  = 1000
    MAX_STEPS = 10000
)

func ReadFile(filepath string) (PipeDiagram, error) {
    pd := make(PipeDiagram, 0, MAX_SIZE)

    f, err := os.Open(filepath)
    if err != nil {
        return nil, errors.New("error opening file")
    }
    defer f.Close()
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        if err != nil {
            panic("error reading file")
        }
        row := make([]PipeType, len(line)-1)
        for i := len(line) - 2; i >= 0; i-- {
            row[i] = ToPipeType(line[i])
        }
        pd = append(pd, row)
    }
    return pd, nil
}

func findEntryPoint(pd *PipeDiagram) Position {
    for i, row := range *pd {
        for j, v := range row {
            if v == EntryPoint {
                return Position{i, j}
            }
        }
    }
    panic("no entry point found")
}

func findFurthestPoint(pd *PipeDiagram, log *LogSteps, entry Position) int {
    visited := make([][]bool, len(*pd))

    for i := range visited {
        visited[i] = make([]bool, len((*pd)[0]))
    }
    maxCount := 0
    var try func(Position, int)
    try = func(pos Position, count int) {
        log.Save(count, pos)
        for _, dir := range []Direction{North, South, East, West} {
            nextPos := pos.NextPos(dir)
            if pd.OutOfBounds(nextPos) || !pos.IsConnected(pd, nextPos) {
                continue
            }
            if nextPos.SameAs(entry) {
                maxCount = max(maxCount, count+1)
                // fmt.Println("MaxCount: ", maxCount, "Path: ", log.String())
            } else if !visited[nextPos.Row][nextPos.Col] {
                visited[nextPos.Row][nextPos.Col] = true
                try(nextPos, count+1)
                visited[nextPos.Row][nextPos.Col] = false
            }
        }
    }
    try(entry, 0)
    return maxCount / 2
}

func countInnerTiles(pd *PipeDiagram) int {
    expanded := pd.Expand()
    for i := 0; i < expanded.Height(); i++ {
        if !expanded.IsBlocked(Position{i, 0}) {
            expanded.Flood(Position{i, 0})
        }
        if !expanded.IsBlocked(Position{i, expanded.Width() - 1}) {
            expanded.Flood(Position{i, expanded.Width() - 1})
        }
    }
    for j := 0; j < expanded.Width(); j++ {
        if !expanded.IsBlocked(Position{0, j}) {
            expanded.Flood(Position{0, j})
        }
        if !expanded.IsBlocked(Position{expanded.Height() - 1, j}) {
            expanded.Flood(Position{expanded.Height() - 1, j})
        }
    }
    fmt.Println(expanded.String())
    return expanded.Count(Ground)
}

func main() {
    if len(os.Args) < 2 {
        panic("args: missing input file path")
    }
    filepath := os.Args[1]
    pd, err := ReadFile(filepath)
    if err != nil {
        panic(err)
    }
    steps := LogSteps{
        Steps: make([]Position, 0, MAX_STEPS),
        N:     0,
    }
    entry := findEntryPoint(&pd)
    fmt.Printf("Part 1: %d\n", findFurthestPoint(&pd, &steps, entry))
    pd.Simplify(steps)
    fmt.Printf("Part 2: %d\n", countInnerTiles(&pd))
}
