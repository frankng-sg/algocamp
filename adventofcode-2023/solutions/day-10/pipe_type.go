package main

type PipeType int

const (
    Error PipeType = iota
    EntryPoint
    Ground
    NorthSouth
    EastWest
    NorthEast
    NorthWest
    SouthWest
    SouthEast
    Empty
    Flooded
)

func ToPipeType(v byte) PipeType {
    switch v {
    case '.':
        return Ground
    case '|':
        return NorthSouth
    case '-':
        return EastWest
    case 'L':
        return NorthEast
    case 'J':
        return NorthWest
    case '7':
        return SouthWest
    case 'F':
        return SouthEast
    case 'S':
        return EntryPoint
    case 'O':
        return Empty
    default:
        return Error
    }
}

func (pt PipeType) Byte() byte {
    switch pt {
    case Ground:
        return '.'
    case NorthSouth:
        return '|'
    case EastWest:
        return '-'
    case NorthEast:
        return 'L'
    case NorthWest:
        return 'J'
    case SouthWest:
        return '7'
    case SouthEast:
        return 'F'
    case EntryPoint:
        return 'S'
    case Flooded:
        return 'O'
    case Empty:
        return ' '
    default:
        return '?'
    }
}
