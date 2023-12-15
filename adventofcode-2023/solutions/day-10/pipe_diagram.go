package main

import "strings"

type PipeDiagram [][]PipeType

func (pd *PipeDiagram) String() string {
	var str strings.Builder

	for _, row := range *pd {
		for _, v := range row {
			str.WriteByte(v.Byte())
		}
		str.WriteString("\n")
	}
	return str.String()
}

func (pd *PipeDiagram) GetType(pos Position) PipeType {
	return (*pd)[pos.Row][pos.Col]
}

func (pd *PipeDiagram) OutOfBounds(pos Position) bool {
	return pos.Row < 0 || pos.Row >= len(*pd) || pos.Col < 0 || pos.Col >= len((*pd)[0])
}

func (pd *PipeDiagram) Height() int {
	return len(*pd)
}

func (pd *PipeDiagram) Width() int {
	return len((*pd)[0])
}

func (pd *PipeDiagram) Expand() PipeDiagram {
	var newMap PipeDiagram = make([][]PipeType, pd.Height()*2-1)
	for i := range newMap {
		newMap[i] = make([]PipeType, pd.Width()*2-1)
	}
	for i := 0; i < pd.Height(); i++ {
		for j := 0; j < pd.Width(); j++ {
			newMap[i*2][j*2] = (*pd)[i][j]
		}
	}
	// Expand horizontally
	for i := 0; i < newMap.Height(); i += 2 {
		for j := 1; j < newMap.Width(); j += 2 {
			left := Position{
				Row: i / 2,
				Col: (j - 1) / 2,
			}
			right := Position{
				Row: i / 2,
				Col: (j + 1) / 2,
			}
			if left.IsConnected(pd, right) {
				newMap[i][j] = EastWest
			} else {
				newMap[i][j] = Empty
			}
		}
	}
	// Expand vertically
	for i := 1; i < newMap.Height(); i += 2 {
		for j := 0; j < newMap.Width(); j += 2 {
			top := Position{
				Row: (i - 1) / 2,
				Col: j / 2,
			}
			bottom := Position{
				Row: (i + 1) / 2,
				Col: j / 2,
			}
			if top.IsConnected(pd, bottom) {
				newMap[i][j] = NorthSouth
			} else {
				newMap[i][j] = Empty
			}
		}
	}
	// Fill vacuum
	for i := 1; i < newMap.Height(); i += 2 {
		for j := 1; j < newMap.Width(); j += 2 {
			newMap[i][j] = Empty
		}
	}
	return newMap
}

func (pd *PipeDiagram) IsBlocked(pos Position) bool {
	t := pd.GetType(pos)
	return t != Empty && t != Ground
}

func (pd *PipeDiagram) Flood(pos Position) {
	var try func(Position)
	try = func(pos Position) {
		if pd.OutOfBounds(pos) || pd.IsBlocked(pos) {
			return
		}
		(*pd)[pos.Row][pos.Col] = Flooded
		for _, dir := range []Direction{North, South, East, West} {
			try(pos.NextPos(dir))
		}
	}
	try(pos)
}

func (pd *PipeDiagram) Count(pipeType PipeType) int {
	count := 0
	for i := 0; i < pd.Height(); i++ {
		for j := 0; j < pd.Width(); j++ {
			if (*pd)[i][j] == pipeType {
				count++
			}
		}
	}
	return count
}

func (pd *PipeDiagram) Reset() {
	for i := 0; i < pd.Height(); i++ {
		for j := 0; j < pd.Width(); j++ {
			(*pd)[i][j] = Ground
		}
	}
}

func (pd *PipeDiagram) Simplify(steps LogSteps) {
	pipeType := make([]PipeType, len(steps.Steps))
	for i := range pipeType {
		loc := steps.Steps[i]
		pipeType[i] = (*pd)[loc.Row][loc.Col]
	}
	pd.Reset()
	for i, loc := range steps.Steps {
		(*pd)[loc.Row][loc.Col] = pipeType[i]
	}
}
