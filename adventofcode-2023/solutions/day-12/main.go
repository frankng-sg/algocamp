package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Input struct {
	Val string
	Err error
}

type RecordType byte

const (
	DAMAGED     RecordType = '#'
	OPERATIONAL RecordType = '.'
	UNKNOWN     RecordType = '?'
)

func readInput(filepath string) chan Input {
	ch := make(chan Input)
	go func() {
		defer close(ch)

		f, err := os.Open(filepath)
		if err != nil {
			ch <- Input{"", err}
			return
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				ch <- Input{"", err}
				return
			}
			line = strings.TrimRight(line, "\n")
			line = strings.TrimRight(line, "\r")
			if line == "" {
				break
			}
			select {
			case ch <- Input{line, nil}:
				//case <-time.After(1 * time.Second):
				//	break
			}
		}
	}()
	return ch
}

func solvePart1(filepath string, solverFn func([]RecordType, []int) int, verbose bool) {
	result := 0
	i := 0
	fmt.Printf("\nPart 1:\n------------------------\n")
	for line := range readInput(filepath) {
		i++
		begin := time.Now()
		if line.Err != nil {
			fmt.Println(line.Err)
			return
		}
		record, reqs := decode(line.Val)
		cnt := solverFn(record, reqs)
		result += cnt
		if verbose {
			fmt.Printf("Line=%d, Count=%d, Time=%s\n", i, cnt, time.Since(begin))
		}
	}
	fmt.Printf("\nAnswer: %d\n", result)
}

func solvePart2(filepath string, solverFn func([]RecordType, []int) int, verbose bool) {
	result := 0
	i := 0
	fmt.Printf("\nPart 2:\n------------------------\n")
	for line := range readInput(filepath) {
		begin := time.Now()
		i += 1
		record, reqs := decode(line.Val)
		cnt := solverFn(expandRecord(record, UNKNOWN, 5), expandReqs(reqs, 5))
		result += cnt
		if verbose {
			fmt.Printf("Line=%d, Count=%d, Time=%s\n", i, cnt, time.Since(begin))
		}
	}

	fmt.Printf("\nAnswer: %d\n", result)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./main <input_file>")
		return
	}
	filepath := os.Args[1]
	solvePart1(filepath, countCombo4, false)
	solvePart2(filepath, countCombo4, false)
}
