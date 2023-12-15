package main

import (
    "strings"
)

type Space byte

const (
    Empty  Space = '.'
    Galaxy       = '#'
)

type Universe [][]Space

func (u *Universe) String() string {
    var str strings.Builder

    for i := range *u {
        for _, v := range (*u)[i] {
            str.WriteByte(byte(v))
        }
        str.WriteString("\n")
    }
    return str.String()
}

func (u *Universe) Height() int {
    return len(*u)
}

func (u *Universe) Width() int {
    return len((*u)[0])
}

func (u *Universe) ShiftDown(rowIndex int) {
    *u = append(*u, make([]Space, u.Width()))

    for i := u.Height() - 1; i > rowIndex; i-- {
        copy((*u)[i], (*u)[i-1])
    }
}

func (u *Universe) ShiftRight(colIndex int) {
    for i := range *u {
        (*u)[i] = append((*u)[i], Empty)
    }

    for j := u.Width() - 1; j > colIndex; j-- {
        for i := range *u {
            (*u)[i][j] = (*u)[i][j-1]
        }
    }
}

func (u *Universe) IsEmptyCol(colIndex int) bool {
    for i := range *u {
        if (*u)[i][colIndex] != Empty {
            return false
        }
    }
    return true
}

func (u *Universe) IsEmptyRow(rowIndex int) bool {
    for _, v := range (*u)[rowIndex] {
        if v != Empty {
            return false
        }
    }
    return true
}

func (u *Universe) DuplicateRow(rowIndex int) {
    u.ShiftDown(rowIndex)
}

func (u *Universe) DuplicateCol(colIndex int) {
    u.ShiftRight(colIndex)
}

func (u *Universe) Expand() {
    i := 0
    for i < u.Height() {
        if u.IsEmptyRow(i) {
            u.DuplicateRow(i)
            i += 2
        } else {
            i += 1
        }
    }
    j := 0
    for j < u.Width() {
        if u.IsEmptyCol(j) {
            u.DuplicateCol(j)
            j += 2
        } else {
            j += 1
        }
    }
}

func (u *Universe) FindEmptyLine(size int, fn func(int) bool) []int {
    l := make([]int, 0, 100)
    for i := 0; i < size; i++ {
        if fn(i) {
            l = append(l, i)
        }
    }
    return l
}

type EmptyLines []int

func (e *EmptyLines) Count(from, to int) int {
    if from > to {
        from, to = to, from
    }
    count := 0
    for _, v := range *e {
        if v > from && v < to {
            count++
        }
    }
    return count
}

func (u *Universe) FindEmptyRows() EmptyLines {
    return u.FindEmptyLine(u.Height(), u.IsEmptyRow)
}

func (u *Universe) FindEmptyCols() EmptyLines {
    return u.FindEmptyLine(u.Width(), u.IsEmptyCol)
}

func (u *Universe) Find(t Space) []Position {
    nodes := make([]Position, 0, 100)
    for i := 0; i < u.Height(); i++ {
        for j := 0; j < u.Width(); j++ {
            if (*u)[i][j] == t {
                nodes = append(nodes, Position{i, j})
            }
        }
    }
    return nodes
}
