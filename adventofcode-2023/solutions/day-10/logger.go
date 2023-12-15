package main

import (
	"fmt"
	"strings"
)

type LogSteps struct {
	Steps []Position
	N     int
}

func (log *LogSteps) Save(idx int, pos Position) {
	if idx >= len(log.Steps) {
		log.Steps = append(log.Steps, pos)
	} else {
		log.Steps[idx] = pos
	}
	log.N = idx + 1
}

func (log *LogSteps) String() string {
	var str strings.Builder
	for i := 0; i < log.N; i++ {
		str.WriteString(fmt.Sprintf("%v ", log.Steps[i]))
	}
	return str.String()
}
