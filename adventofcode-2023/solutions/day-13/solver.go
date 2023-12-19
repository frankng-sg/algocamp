package main

func solve(pattern *Pattern) int {
    line, isVertical := findMirror(pattern)
    if isVertical {
        return line
    }
    return line * 100
}

func findMirror(pattern *Pattern) (line int, isVertical bool) {
    height := len(*pattern)
    width := len((*pattern)[0])
    for i := 1; i < height; i++ {
        if isMirror(pattern, i, false) {
            return i, false
        }
    }
    for i := 1; i < width; i++ {
        if isMirror(pattern, i, true) {
            return i, true
        }
    }
    return 0, false
}

func isEqual(pattern *Pattern, idx1, idx2 int, isVertical bool) bool {
    if isVertical {
        for i := 0; i < len(*pattern); i++ {
            if (*pattern)[i][idx1] != (*pattern)[i][idx2] {
                return false
            }
        }
    } else {
        for j := 0; j < len((*pattern)[0]); j++ {
            if (*pattern)[idx1][j] != (*pattern)[idx2][j] {
                return false
            }
        }
    }
    return true
}

func isMirror(pattern *Pattern, line int, isVertical bool) bool {
    i, j := line-1, line
    size := len(*pattern)
    if isVertical {
        size = len((*pattern)[0])
    }
    for i >= 0 && j < size {
        if !isEqual(pattern, i, j, isVertical) {
            return false
        }
        i--
        j++
    }
    return true
}
