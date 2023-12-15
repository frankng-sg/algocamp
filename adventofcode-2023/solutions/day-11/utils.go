package main

func ToSpace(b byte) Space {
	switch b {
	case '.':
		return Empty
	case '#':
		return Galaxy
	default:
		panic("unknown value")
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Distance(a, b Position) int {
	return abs(a.Row-b.Row) + abs(a.Col-b.Col)
}
