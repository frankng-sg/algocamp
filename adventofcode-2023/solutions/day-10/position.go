package main

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type Position struct {
	Row, Col int
}

func (p Position) SameAs(other Position) bool {
	return p.Row == other.Row && p.Col == other.Col
}

func (p Position) GetLinkedPos(pipeType PipeType) []Position {
	switch pipeType {
	case NorthSouth:
		return []Position{p.NextPos(North), p.NextPos(South)}
	case EastWest:
		return []Position{p.NextPos(East), p.NextPos(West)}
	case NorthEast:
		return []Position{p.NextPos(North), p.NextPos(East)}
	case NorthWest:
		return []Position{p.NextPos(North), p.NextPos(West)}
	case SouthWest:
		return []Position{p.NextPos(South), p.NextPos(West)}
	case SouthEast:
		return []Position{p.NextPos(South), p.NextPos(East)}
	case EntryPoint:
		return []Position{p.NextPos(North), p.NextPos(South), p.NextPos(East), p.NextPos(West)}
	default:
		return []Position{}
	}
}

func (p Position) IsConnected(pd *PipeDiagram, other Position) bool {
	reachable := make(map[Position]bool)
	for _, pos := range p.GetLinkedPos(pd.GetType(p)) {
		reachable[pos] = true
	}
	for _, pos := range other.GetLinkedPos(pd.GetType(other)) {
		reachable[pos] = true
	}
	if _, ok := reachable[p]; !ok {
		return false
	}
	if _, ok := reachable[other]; !ok {
		return false
	}
	return true
}

func (p Position) NextPos(dir Direction) Position {
	switch dir {
	case North:
		return Position{p.Row - 1, p.Col}
	case South:
		return Position{p.Row + 1, p.Col}
	case East:
		return Position{p.Row, p.Col + 1}
	case West:
		return Position{p.Row, p.Col - 1}
	default:
		panic("invalid direction")
	}
}
